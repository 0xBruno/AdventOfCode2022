package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	var highestCalories int = 0
	var caloricTracker int = 0

	for _, line := range fileLines {

		calories, err := strconv.Atoi(line)

		if err != nil {
			// set the highestCalories to the tracker if it is greater
			if caloricTracker > highestCalories {
				highestCalories = caloricTracker
			}

			// reset the caloricTracker
			caloricTracker = 0
		}

		caloricTracker += calories
	}

	fmt.Println(highestCalories)
}
