package main

import "fmt"

// struct
type me struct {
	name  string
	age   int
	place string
}

// func taking stuct as param

func (m me) intro() {
	fmt.Println("Hi ", m.name)
	fmt.Println("You're age is : ", m.age)
	fmt.Println("you're from ", m.place)
}

func (m *me) bday() {
	fmt.Println("Happy Birthday now you're age is ")
	m.age++
}

func main() {

	m1 := me{
		name:  "harsha",
		age:   21,
		place: "TPT"}

	m1.intro()
	m1.bday()

	m1.intro()

	// temp struct for single time use
	other := struct {
		name string
		age  int
	}{"batman", 29}

	fmt.Printf("you're %s and youre age is %d", other.name, other.age)
}
