package main

import (
	"errors"
	"fmt"
)

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("zero is not allowed")
	}

	return a / b, nil
}

func main() {
	res, err := div(10, 2)
	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println(res)
	}

	res2, err2 := div(10, 0)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(res2)
	}

}
