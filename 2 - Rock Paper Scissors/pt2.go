package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// column 1
// A for rock
// B for paper
// C for scissors

// column 2
// X issa l
// Y issa tie
// Z issa w

// score for single round
// 1 for rock
// 2 for paper
// 3 for scissors

// PLUS outcome of the round
// 0 if you lost
// 3 if the round was a draw
// 6 if you won

func main() {
	// strategy guide map
	var inputMap = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	readFile, err := os.Open(os.Args[1])
	defer readFile.Close()

	if err != nil {
		log.Fatalf("Unable to open program input file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var totalScore int
	for fileScanner.Scan() {
		var roundScore int

		choices := strings.Fields(fileScanner.Text())

		elfChoice := inputMap[choices[0]]
		myChoice := inputMap[choices[1]]

		// rock paper scissors
		if myChoice == 2 {
			// it's a tie
			roundScore = 3 + elfChoice

		} else if myChoice == 3 {
			// i win
			var choicePoints int
			if elfChoice == 1 {
				choicePoints = 2
			} else if elfChoice == 2 {
				choicePoints = 3
			} else {
				choicePoints = 1
			}
			roundScore = 6 + choicePoints

		} else {
			// elf wins
			var choicePoints int
			if elfChoice == 1 {
				choicePoints = 3
			} else if elfChoice == 2 {
				choicePoints = 1
			} else {
				choicePoints = 2
			}
			roundScore = 0 + choicePoints
		}

		totalScore += roundScore

	}

	fmt.Println(totalScore)

}
