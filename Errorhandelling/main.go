package main

import (
	"fmt"
	"strconv"
)

func main() {
	numstr := "abc"
	num, err := strconv.Atoi(numstr)

	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("num:", num)

}
