package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	// points map
	pointsMap := make(map[string]int)
	var lowerPoints int = 0
	var upperPoints int = 26

	for i := 'a'; i <= 'z'; i++ {
		I := unicode.ToUpper(i)

		pointsMap[string(i)] = lowerPoints + 1
		pointsMap[string(I)] = upperPoints + 1

		lowerPoints += 1
		upperPoints += 1
	}

	// read lines
	readFile, err := os.Open(os.Args[1])
	defer readFile.Close()

	if err != nil {
		log.Fatalf("Unable to open program input file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var elfGroups [][]string
	var group []string

	for fileScanner.Scan() {

		line := fileScanner.Text()
		if len(group) == 3 {
			elfGroups = append(elfGroups, group)
			group = nil
		}

		group = append(group, line)

	}
	// add the last group from the loop
	// to the slice
	elfGroups = append(elfGroups, group)

	var totalScore int

	for _, ruckGroup := range elfGroups {
		freqTracker := make(map[string]int)

		fmt.Println("[RUCKGROUP]")
		for _, ruck := range ruckGroup {
			itemTracker := make(map[string]int32)
			fmt.Println("[RUCK] ", ruck)
			for _, item := range ruck {
				// Count the item
				itemTracker[string(item)] = 1

			}

			for item, _ := range itemTracker {
				freqTracker[item] += 1
			}

		}

		for item, occurrences := range freqTracker {
			// does the item show up in all rucks?
			if occurrences == 3 {
				totalScore += pointsMap[item]
			}
		}

	}

	fmt.Println(totalScore)

}
