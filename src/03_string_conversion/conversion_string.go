package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		origin string = "ABC"
		an     int
		newS   string
	)
	fmt.Println(strconv.IntSize)
	an, err := strconv.Atoi(origin)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting width error\n", origin)
		return
	}
	fmt.Println("an", an)
	an += 5
	newS = strconv.Itoa(an)
	fmt.Println(newS)
}
