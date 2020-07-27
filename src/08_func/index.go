package main

import (
	"errors"
	"fmt"
	"math"
)

var num int = 10
var numx2, numx3, sum, multiply, poor int

func main() {

	numx2, numx3 = getX2AndX3(num)
	fmt.Println(num, numx2, numx3)

	numx2, numx3 = getX2AndX3_2(num)
	fmt.Println(num, numx2, numx3)

	sum, multiply, poor = multReturnval(num, num)
	fmt.Println(sum, multiply, poor)

	sum, multiply, poor = multReturnval2(num, num)
	fmt.Println(sum, multiply, poor)

	if ret1, err1 := errorReturnVal(4); err1 != nil {
		fmt.Println("error", ret1, err1)
	} else {
		fmt.Println("ok", ret1, err1)
	}

	if ret2, err2 := errorReturnVal2(-4); err2 != nil {
		fmt.Println("error", ret2, err2)
	} else {
		fmt.Println("ok", ret2, err2)
	}

	var val int = 10
	sideEffect(num, num, &val)
	fmt.Println("val", val)

	slice := []int{1, 2, 3, 4, 6}
	fmt.Println("slice", slice)
	minVal := min(slice...)
	fmt.Println("minVal", minVal)
}

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	return
}

// 非命名返回值
func multReturnval(input1, input2 int) (int, int, int) {
	return input1 + input2, input1 * input2, input1 - input2
}

// 命名返回值
func multReturnval2(input1, input2 int) (sum int, multiply int, poor int) {
	sum, multiply, poor = input1+input2, input1*input2, input1-input2
	return
}

// 非命名返回值
func errorReturnVal(val float64) (float64, error) {
	if val < 0 {
		return float64(math.NaN()), errors.New("error")
	}
	return math.Sqrt(val), nil
}

// 命名返回值
func errorReturnVal2(val float64) (ret float64, err error) {
	if val < 0 {
		ret = float64(math.NaN())
		err = errors.New("error")
	} else {
		ret = math.Sqrt(val)
	}
	return
}

// 改变外部变量
func sideEffect(i, j int, reply *int) {
	*reply = i * j
}

// 传递变长参数
func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	var min int = s[0]
	for _, val := range s {
		if min > val {
			min = val
		}
	}
	return min
}
