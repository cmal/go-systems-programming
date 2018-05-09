package main

import (
	"fmt"
	"os"
	"strconv"
	"math"
)

func main() {
	var max, min int

	max = math.MinInt32
	min = math.MaxInt32

	arguments := os.Args

	for i := 1; i < len(arguments); i ++ {
		s := arguments[i]
		val, _ := strconv.Atoi(s)
		if max < val {
			max = val
		}
		if min > val {
			min = val
		}
	}
	fmt.Println("max: ", max, " min: ", min)
}
