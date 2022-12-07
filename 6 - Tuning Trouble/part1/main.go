package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	dataStream, err := os.ReadFile("./input")

	if err != nil {
		log.Fatalf("Could not read input file! %s\n", err)
	}

	for i, _ := range dataStream {

		// bounds check
		if i >= len(dataStream)-14 {
			break
		}

		chunk := dataStream[i : i+14]
		repeatsFlag := false
		startOfPacketMarker := i + 14 // is inclusive of the chunk

		for _, i := range chunk {
			if strings.Count(string(chunk), string(i)) > 1 {
				repeatsFlag = true
			}
		}

		if repeatsFlag == false {
			fmt.Printf("Start of packet marker is at position: %d\n", startOfPacketMarker)
			break
		}
	}
}
