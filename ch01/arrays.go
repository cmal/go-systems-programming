package main

import "fmt"

func main() {

	myArray := [4]int{1,2,4,-4}
	twoD := [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	threeD := [2][2][2]int{{{1,2},{3,4}},{{5,6},{7,8}}}

	myArray[0] = 3
	twoD[1][2] = 15
	threeD[0][1][1] = -1

	fmt.Println(myArray[-1], myArray[10])
	fmt.Println(twoD) // when pass an array to a function, you actually pass a copy of the array
	fmt.Println(threeD[-1][20][0])
}

