package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {

	rand.Seed(37)
	mySlice := make([]int, 0)
	mySlice = append(mySlice, rand.Intn(100))
	mySlice = append(mySlice, rand.Intn(100))
	mySlice = append(mySlice, rand.Intn(100))
	mySlice = append(mySlice, rand.Intn(100))
	mySlice = append(mySlice, rand.Intn(100))
	mySlice = append(mySlice, rand.Intn(100))
	mySlice = append(mySlice, rand.Intn(100))
	mySlice = append(mySlice, rand.Intn(100))
	fmt.Println("=:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i] < mySlice[j]
	})
	fmt.Println("<:", mySlice)
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i] > mySlice[j]
	})
	fmt.Println(">:", mySlice)

}
