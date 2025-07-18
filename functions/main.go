package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func area(length, width int) (area int) {
	area = length * width
	return
}

func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	fmt.Println("Add:", add(5, 3))

	q, r := divide(10, 3)
	fmt.Println("Divide: Quotient =", q, ", Remainder =", r)

	a := area(4, 6)
	fmt.Println("Area:", a)

	total := sumAll(1, 2, 3, 4, 5)
	fmt.Println("Sum:", total)
}
