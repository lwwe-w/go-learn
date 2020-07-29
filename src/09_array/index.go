package main

import (
	"fmt"
)

func main() {
	var arr1 = new([5]int) // &[0,0,0,0,0] 指针地址

	var arr [5]int // [0,0,0,0,0] 值

	runFor(arr)

	runRange(arr)

	runString()

	runArrayValue()

	runFibonacciArray()

	runPointerArray()

	multidimArray()

	fmt.Println(arr, arr1)
}

func runFor(arr [5]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(i)
	}
}

func runRange(arr [5]int) {
	for _, i := range arr {
		fmt.Println(i)
	}
}

func runString() {
	b := [...]string{"a", "b", "c", "d"}
	for i := range b {
		fmt.Println(i, b[i])
	}
}

func runArrayValue() {
	var arr [2]int
	for _, v := range arr {
		v = 4
		fmt.Println("v", v)
	}
	fmt.Println("runArrayValue", arr)
}

func runFibonacciArray() {
	var (
		arr [50]int
	)
	for k, v := range arr {
		if k <= 1 {
			arr[k] = 1
		} else {
			arr[k] = arr[k-2] + arr[k-1]
		}
		fmt.Println("v", v)
	}
	fmt.Println("runFibonacciArray", arr)
}

func runPointerArray() {
	for i := 0; i < 3; i++ {
		fp(&[...]int{i, i * i, i * i * i})
	}
}

func fp(a *[3]int) { fmt.Println(a) }

func multidimArray() {
	const (
		WIDTH  = 20
		HEIGHT = 10
	)
	type pixel int
	var screen [WIDTH][HEIGHT]pixel

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
		}
	}

	fmt.Println("multidimArray", screen)
}
