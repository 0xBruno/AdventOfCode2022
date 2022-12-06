package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Instruction struct {
	amount int
	from   int
	to     int
}

func NewInstruction(amount int, from int, to int) *Instruction {
	return &Instruction{
		amount: amount,
		from:   from,
		to:     to,
	}
}

type SupplyStack struct {
	id    int
	stack []rune
}

func (supplyStack *SupplyStack) Peek() string {
	// get last item from stack
	return string(supplyStack.stack[len(supplyStack.stack)-1])
}

func NewSupplyStack(id int, stack []rune) *SupplyStack {
	return &SupplyStack{
		id:    id,
		stack: stack,
	}
}

type CargoCrane struct {
	ops int
}

func NewCargoCrane() *CargoCrane {
	return &CargoCrane{ops: 0}
}

func (crane *CargoCrane) Rearrange(stacks []SupplyStack, instructions []Instruction) []SupplyStack {
	/*
		STACKMAP structure
		{ <stack> : <stack>}

		basic algo pseudocode.
		With the stackMap, lookups for a specific stack are easier,
		then just get the amount and append it to the instruction.to
		one at a time

		stackMap[instruction.from]
		stackMap[instruction.from][instruction.amount]
		stackMap[instruction.to] := append(stackMap[instruction.to], amount)
	*/

	// 1. create the internal stackMap
	stackMap := make(map[int][]rune)
	for _, stack := range stacks {
		stackMap[stack.id] = stack.stack
	}

	// 2. execute the instructions
	var fromStack []rune
	var toStack []rune

	for _, instruction := range instructions {

		fromStack = stackMap[instruction.from]
		toStack = stackMap[instruction.to]

		for i := instruction.amount; i > 0; i-- {
			crane.ops += 1
			var poppedCrate rune
			fmt.Printf("Executing instruction: move %d from %d to %d\n",
				instruction.amount,
				instruction.from,
				instruction.to)

			// pop the crate from the fromStack
			// gets the poppedCrate and reduces the fromStack by 1 essentially deleting it
			poppedCrate, fromStack = fromStack[len(fromStack)-1], fromStack[:len(fromStack)-1]

			// push the crate to the toStack
			toStack = append(toStack, poppedCrate)

		}

		stackMap[instruction.from] = fromStack
		stackMap[instruction.to] = toStack

	}

	var organizedStacks []SupplyStack

	for stackId, stack := range stackMap {
		tmpStack := NewSupplyStack(stackId, stack)
		organizedStacks = append(organizedStacks, *tmpStack)
	}

	return organizedStacks
}

func main() {
	// read lines
	readFile, err := os.Open("./input")
	defer readFile.Close()

	if err != nil {
		log.Fatalf("Unable to open program input file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var rawData [][]rune

	for fileScanner.Scan() {
		line := fileScanner.Text()
		rawData = append(rawData, []rune(line))
	}

	var stackHeader []rune
	var stacks [][]rune
	var instructions []Instruction

	stackHeader, stacks, instructions, err = parseInput(rawData)

	if err != nil {
		log.Fatal(err)
	}

	supplyStacks, err := makeSupplyStacks(stackHeader, stacks)

	if err != nil {
		log.Fatalf("ERROR in makeSupplyStacks: %s", err)
	}

	var crane *CargoCrane = NewCargoCrane()

	organizedStacks := crane.Rearrange(supplyStacks, instructions)
	fmt.Println("[*] Stacks are reorganized!")
	fmt.Printf("[*] %d operations performed!\n", crane.ops)

	topOfStacks(organizedStacks)

}

func topOfStacks(supplyStacks []SupplyStack) {

	// since the []SupplyStack param is not
	// guaranteed to be sorted, we need to order first
	sortMap := make(map[int]string)
	for _, stack := range supplyStacks {
		sortMap[stack.id] = stack.Peek()
	}

	// so many for loops....
	keySlice := make([]int, 0)
	for key, _ := range sortMap {
		keySlice = append(keySlice, key)
	}

	//sort the slice
	sort.Ints(keySlice)
	var output string
	for _, key := range keySlice {
		output += sortMap[key]
	}

	fmt.Println("The new top of each stack is: ", output)
}

func parseInput(rawData [][]rune) ([]rune, [][]rune, []Instruction, error) {
	var stackData [][]rune
	var numOfStacks []rune
	var instructions []Instruction

	fmt.Print("[*] Reading in supply stacks data...\r\n")
	fmt.Print("[*] Reading instructions...\r\n")

	for _, runeList := range rawData {

		// ignore the divider
		if len(runeList) == 0 {
			continue
		}

		switch firstChar := runeList[0]; firstChar {
		// matches the '[' rune
		case 91:
			// parse the data
			// e.g. '[P] [G] [R]' and not '1 2 3 4' also not 'move x from y to z'
			stackData = append(stackData, runeList)
		// the letter 'm' e.g. move x from y to z
		case 109:
			// parse the instructions
			instructString := string(runeList)
			instructSlice := strings.Fields(instructString)

			amount, err := strconv.Atoi(instructSlice[1])
			from, err := strconv.Atoi(instructSlice[3])

			to, err := strconv.Atoi(instructSlice[5])

			if err != nil {
				return nil, nil, nil, fmt.Errorf("[!]ERR: Unable to parse instructions from input %s", err)
			}

			instruction := NewInstruction(amount, from, to)
			instructions = append(instructions, *instruction)
		case 32:
			// firstChar is space '32' code point and next is 1 '49' code point
			// e.g. 1 2 3 4 5
			if runeList[1] != 49 {
				continue
			}
			numOfStacks = runeList
		}

	}

	return numOfStacks, stackData, instructions, nil
}

func makeSupplyStacks(stackHeader []rune, stacks [][]rune) ([]SupplyStack, error) {
	fmt.Println("[*] Reading number of supply stacks...")
	var supplyStacks []SupplyStack
	for i, columnNum := range stackHeader {
		// 32 is the code point for a space
		// so we ignore
		if columnNum == 32 {
			continue
		}

		var tempStack []rune
		for _, stack := range stacks {
			//bounds checking
			if i >= len(stack) {
				continue
			}

			// dont add the spaces to the stacks
			if stack[i] == 32 {
				continue
			}

			tempStack = append(tempStack, stack[i])
		}
		fmt.Printf("[*] Supply stack #%s:%s\n", string(columnNum), string(tempStack))

		stackId, err := strconv.Atoi(string(columnNum))

		if err != nil {
			return nil, err
		}

		// reverse the stack so that the ordering is correct
		tempStack = reverse(tempStack)

		stack := NewSupplyStack(stackId, tempStack)
		supplyStacks = append(supplyStacks, *stack)
	}

	return supplyStacks, nil
}

// reverse a slice
// https://go.dev/play/p/vkJg_D1yUb
func reverse(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
