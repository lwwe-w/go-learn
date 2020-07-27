package main

import (
	"fmt"
)

func main() {
	var a int = 100
	switch a {
	case 99:
		fmt.Println("It's equal to 98")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("default")
	}

	switch {
	case a < 0:
		fmt.Println("Number is negative")
	case a > 0 && a < 100:
		fmt.Println("Number is between 0 and 100")
	default:
		fmt.Println("Number is 100 or greater")
	}

	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
