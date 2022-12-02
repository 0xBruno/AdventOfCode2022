package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	readFile, err := os.Open(os.Args[1])
	defer readFile.Close()

	if err != nil {
		log.Fatalf("Unable to open program input file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	var allCalories []int
	var caloricTracker int = 0

	for _, line := range fileLines {

		calories, err := strconv.Atoi(line)

		if err != nil {
			// add total calories per elf to allCalories slice
			allCalories = append(allCalories, caloricTracker)

			// reset the caloricTracker
			caloricTracker = 0
		}

		caloricTracker += calories
	}

	// sort
	sort.Ints(allCalories)

	var result int
	// add last 3 elements in slice to result
	for _, cals := range allCalories[len(allCalories)-3:] {
		result += cals
	}

	fmt.Println(result)

}
