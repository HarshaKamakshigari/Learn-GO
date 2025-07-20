package main

import "fmt"

func main() {

	var nums []int
	fmt.Println(nums)

	t := []int{1, 2, 3, 4, 5}
	fmt.Println(t)

	s := make([]string, 3)
	s[0] = "a"
	fmt.Println(s)

	s = append(s, "b", "c")
	fmt.Println(s)

	c := make([]string, 3)

	copy(c, s)
	fmt.Println(c)

}
