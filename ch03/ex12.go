package main

import (
	"fmt"
	"strings"
	"os"
	"strconv"
	"log"
)

func main() {

	var s [3]string
	s[0] = "1 b 3 1 a a b"
	s[1] = "11 a 1 1 1 1 a a"
	s[2] = "-1 b 1 -4 a 1"

	var freq int

	if (len(os.Args) < 2) {
		log.Fatal("A minimum frequency should be given as the first command line arg")
	}
	
	freq, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatal("Frequency should be a number.")
	}

	counts := make(map[string]int)

	for i := 0; i < len(s); i++ {
		data := strings.Fields(s[i])
		for _, word := range data {
			_, ok := counts[word]
			if ok {
				counts[word] = counts[word] + 1
			} else {
				counts[word] = 1
			}
		}
	}

	for key, _ := range counts {
		cnt := counts[key]
		if cnt >= freq {
			fmt.Printf("%s -> %d \n", key, cnt)
		}
	}
}
