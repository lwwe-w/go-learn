package main

import (
	"fmt"
)

func main() {
	// int
	i := 5
	fmt.Println(&i)
	var intP *int
	intP = &i
	fmt.Println(intP, *intP)

	// string
	s := "world"
	var p *string = &s
	*p = "d"
	fmt.Println(&s, p, s)

	// 自动反向引用 对空指针的反向引用是不合法的 是程序崩溃
	var d *int = nil
	*d = 12
	fmt.Println(d)
}
