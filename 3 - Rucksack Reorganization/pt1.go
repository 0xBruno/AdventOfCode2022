package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

	var totalScore int = 0

	for fileScanner.Scan() {
		line := fileScanner.Text()

		// split in half
		compartment1 := line[:len(line)/2]
		compartment2 := line[len(line)/2:]

		// compare for repeated chars
		charIndex := strings.IndexAny(compartment1, compartment2)

		// matched character unicode point
		matchedChar := compartment1[charIndex]

		// get priority points from map and add to totalScore
		totalScore += pointsMap[string(matchedChar)]
	}

	fmt.Println(totalScore)

}
