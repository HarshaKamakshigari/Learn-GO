package main

import "fmt"

func main() {

	// IF ELSE STATEMENTS
	var age int
	fmt.Println("Enter age:")
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("You are an adult")
	} else if age >= 13 {
		fmt.Println("You are an teenager")
	} else {
		fmt.Println("You are an child")
	}

	//SWITCH STATEMENTS

	var day int
	fmt.Println("Enter(1-7):")
	fmt.Scanln(&day)

	switch day {
	case 1:
		fmt.Println("monday")

	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6, 7:
		fmt.Println("Weekend!")
	default:
		fmt.Println("Invalid day number")
	}

	// FOR LOOP

	var n int
	fmt.Print("Enter a number to count up to: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}
