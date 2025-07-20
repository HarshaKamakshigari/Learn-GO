package main

import "fmt"

func operate(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func multiply(a, b int) int {
	return a * b
}

func main() {
	result := operate(3, 4, multiply)
	fmt.Println(result)
}
