package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i, max, min int
	var s string
	fmt.Scanln(&s)

	i, _ = strconv.Atoi(s)
	max = i
	min = i

	for i != 0 {
		fmt.Scanln(&s)
		i, _ = strconv.Atoi(s)
		if max < i {
			max = i
		}
		if min > i {
			min = i
		}
	}
	fmt.Println("max: ", max, " min: ", min)
}
