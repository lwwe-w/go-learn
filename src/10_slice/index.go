package main

import (
	"fmt"
)

func main() {
	arraySlice()
	replaceArrayWidthSlice()
	makeSlice()
	question1()
	question2()

	slice1 := make([]byte, 2, 4)
	data := []byte{'e', 'e'}
	fmt.Println("append", append(slice1, data))

	fmt.Println(question3())

	fmt.Println("sumArray", sumArray([]float32{2.0, 15.0, 34.2}))
	sum, average := sumAndAverage([]int{1, 2, 3, 2})
	fmt.Println("sumAndAverage", sum, average)

	fmt.Println("min", minSlice([]int{3, 1, 4, 6}))

	copyAppendSlice()

	fmt.Println(question4(make([]int, 3, 20), 4))

	fmt.Println("question5", question5([]int{1, 2, 3, 4}, even))

	s := []int{1, 2, 3, 4}
	insert := []int{-2, -1, 0}
	fmt.Println(insertStringSlice(s, insert, 1))

	fmt.Println(removeStringSlice(s, 1, 2))

	fmt.Println(splitString("hello", 2))

	fmt.Println(reverseString("Google"))

	fmt.Println(mapFunction(product, []int{1, 2, 3}))
}

func arraySlice() {
	var (
		arr1   [6]int
		slice1 []int = arr1[2:5]
	)

	for k, v := range arr1 {
		arr1[k] = k
		fmt.Println(v)
	}

	for k, v := range slice1 {
		fmt.Println(k, v)
	}

	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// 扩展切片
	slice1 = slice1[:cap(slice1)]
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

}

// 用切片代替数据传参
func replaceArrayWidthSlice() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(sum(a[:]))
}

func sum(a []int) (sum int) {
	for _, v := range a {
		sum += v
	}
	return
}

// 使用make生成slice
func makeSlice() {
	var slice1 []int = make([]int, 3, 4)
	for k, v := range slice1 {
		slice1[k] = k
		fmt.Println(v)
	}
	fmt.Println("slice1", slice1)
}

// question
func question1() {
	s := make([]byte, 5)
	s = s[2:4]
	fmt.Println("s", len(s), cap(s))
}

func question2() {
	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	s2[1] = 't'
	fmt.Println(s1, s2)
}

func append(slice, data []byte) []byte {
	slice = slice[:]
	for _, v := range data {
		sliceLen := len(slice) + 1
		if sliceLen <= cap(slice) {
			slice = slice[:sliceLen]
			slice[len(slice)-1] = v
		}
	}
	return slice
}

func question3() [5]int {
	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
	}
	return items
}

func sumArray(arr []float32) (sum float32) {
	for _, v := range arr {
		sum += v
	}
	return
}

func sumAndAverage(arr []int) (int, float32) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum, float32(sum / len(arr))
}

func minSlice(arr []int) (min int) {
	for k, v := range arr {
		if k == 0 {
			min = v
		} else {
			if min > v {
				min = v
			}
		}
	}
	return
}

func maxSlice(arr []int) (max int) {
	for k, v := range arr {
		if k == 0 {
			max = v
		} else {
			if max < v {
				max = v
			}
		}
	}
	return
}

// 切片的复制与追加
func copyAppendSlice() {
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)

	n := copy(slTo, slFrom)
	fmt.Println(slFrom, slTo, n)

	sl3 := []byte{'w'}
	sl4 := []byte{'q'}
	sl3 = append(sl3, sl4)
	fmt.Println(sl3)
}

func question4(s []int, factor int) []int {
	newS := make([]int, len(s)*factor)
	copy(newS, s)
	s = newS
	return s
}

func question5(s []int, fn func(int) bool) []int {
	newS := make([]int, 0, len(s))
	for _, v := range s {
		if fn(v) {
			newS = newS[:len(newS)+1]
			newS[len(newS)-1] = v
		}
	}
	return newS
}

func even(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

// 将切片插入到另一个切片的指定位置
func insertStringSlice(slice, insert []int, idx int) []int {
	res := make([]int, len(slice)+len(insert))
	at := copy(res, slice[:idx])
	at += copy(res[at:], insert)
	copy(res[at:], slice[idx:])
	return res
}

// 将从start到end索引的元素从切片中移除
func removeStringSlice(slice []int, start, end int) []int {
	var p []int = make([]int, len(slice)+start-end)
	at := copy(p, slice[:start])
	copy(p[at:], slice[end:])
	slice = p
	return slice
}

func splitString(s string, idx int) (str1 string, str2 string) {
	bytes := []byte(s)
	str1 = string(bytes[:idx])
	str2 = string(bytes[idx:])
	return
}

func reverseString(s string) (res string) {
	sByte := []byte(s)
	n, h := len(sByte), len(sByte)/2
	for i := 0; i < h; i++ {
		sByte[i], sByte[n-i-1] = sByte[n-i-1], sByte[i]
	}
	res = string(sByte)
	return
}

func mapFunction(fn func(int) int, list []int) (res []int) {
	res = make([]int, len(list))
	for k, v := range list {
		res[k] = fn(v)
	}
	return
}

func product(num int) (res int) {
	res = num * 10
	return
}
