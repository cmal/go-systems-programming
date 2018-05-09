package main

import "fmt"

func a1() {
	for i := 0; i < 3; i ++ {
		defer fmt.Print(i, " ")
	}
}

func a2() {
	for i := 0; i < 3; i ++ {
		// func body is evaluated after the loop ends
		// i then refered to the local var: i = 3
		defer func() { fmt.Print(i, " ") }()
	}
}

func a3() {
	for i := 0; i < 3; i ++ {
		defer func(n int) { fmt.Print(n, " ") }(i)
	}
}

func main() {
	a1()
	fmt.Println()
	a2()
	fmt.Println()
	a3()
	fmt.Println()
}