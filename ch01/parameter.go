package main

import (
	"fmt"
	"os"
	"strings"
)

// see also:
// fmt.Scanln(), fmt.Scanf(), fmt.Sscanf()
// fmt.Sscanln(), fmt.Sscan()

func main() {
	arguments := os.Args
	minusI := false

	for i:=0; i<len(arguments); i++ {
		if strings.Compare(arguments[i], "-i") == 0 {
			minusI = true
			break
		}
	}

	if minusI {
		fmt.Println("Got the -i parameter!")
		fmt.Print("y/n: ")
		var answer string
		fmt.Scanln(&answer) // here use a pointer ref
		fmt.Println("You entered: ", answer)
	} else {
		fmt.Println("The -i parameter is not set")
	}
}

