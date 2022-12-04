package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read lines
	readFile, err := os.Open(os.Args[1])
	defer readFile.Close()

	if err != nil {
		log.Fatalf("Unable to open program input file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var overlapping int
	for fileScanner.Scan() {
		var pair []int

		pairs := strings.Split(fileScanner.Text(), ",")
		for _, sectionAssignments := range pairs {
			bounds := strings.Split(sectionAssignments, "-")

			for _, sectionIDString := range bounds {
				sectionIDInt, err := strconv.Atoi(sectionIDString)

				if err != nil {
					log.Fatalf("Something is very wrong. Attempted to cast non-int string to int %s", err)
				}
				pair = append(pair, sectionIDInt)
			}

		}

		firstLowerBound := pair[0]
		firstUpperBound := pair[1]
		secondLowerBound := pair[2]
		secondUpperBound := pair[3]

		if firstLowerBound <= secondLowerBound &&
			firstUpperBound >= secondUpperBound {
			fmt.Println("The first assignment contains the second")
			fmt.Println(pair)
			overlapping += 1
		} else if secondLowerBound <= firstLowerBound &&
			secondUpperBound >= firstUpperBound {
			fmt.Println("The second assignment contains the first")
			fmt.Println(pair)
			overlapping += 1

		}

	}

	fmt.Println("Number of overlapping assigned sections: ", overlapping)
}
