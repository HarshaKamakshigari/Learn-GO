package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	url := "https://google.com"

	start := time.Now()

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("the error is: ", err)
		return
	}

	defer resp.Body.Close()

	duration := time.Since(start)

	fmt.Println(resp.StatusCode)
	fmt.Println(duration)

}
