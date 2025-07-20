package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	var f func(int, int) int
	f = add
	fmt.Println(f(2, 3))
}
