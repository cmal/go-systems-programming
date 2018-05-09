package main

import "fmt"

func main() {
	small := []int{}
	large := []int{0,0,0,0}

	origin := []int{1,2,3}

	res := copy(small, origin) // this will not copy anything
	fmt.Println(res, small)

	res = copy(large, origin) // large will be [1 2 3 0]
	fmt.Println(res, large)
}
