package main

import (
	"fmt"
)

func main() {
	forLoop()
	forCharacter()
	fizzBuzz()
	rectangleStars()
	for2()
	test1()
	test2()
	label()
	testgoto()
	test3()
	test4()
}

func forLoop() {
	for i := 0; i < 15; i++ {
		fmt.Println(i)
	}
}

func forCharacter() {
	for i := 1; i < 26; i++ {
		result := ""
		for j := 0; j < i; j++ {
			result += "G"
		}
		fmt.Println(result)
	}
}

func fizzBuzz() {
	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		}
	}
}

func rectangleStars() {
	for i := 0; i < 10; i++ {
		result := ""
		for j := 0; j < 20; j++ {
			if i == 0 || i == 9 || j == 0 || j == 19 {
				result += "*"
			} else {
				result += "!"
			}
		}
		fmt.Println(result)
	}
}

func label() {
LAB:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				break LAB
			}
			fmt.Println(i, j)
		}
	}
}

func testgoto() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}

func for2() {
	var i int = 5
	for i >= 0 {
		i--
		fmt.Println(i)
	}
}

func test1() {
	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d", v)
		v = 5
	}
}

func test2() {
	s := ""
	for s != "aaaaa" {
		fmt.Println(s)
		s = s + "a"
	}
}

func test3() {
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Println(i)
		i++
	}
}

func test4() {
	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
