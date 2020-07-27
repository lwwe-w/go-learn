package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = " Hi, I'm Marc, Hi "
	// 获取字符索引
	fmt.Println(strings.Index(str, "Marc"))
	fmt.Println(strings.Index(str, "Burger"))

	// 从最后往前寻找
	fmt.Println(strings.LastIndex(str, "Hi"))

	// 字符串替换 返回新的字符串
	strings.Replace(str, "Hi", "Test", -1)

	// 统计字符串出现次数
	fmt.Println(strings.Count(str, "Hi"))

	// 重复字符串
	fmt.Println(strings.Repeat(str, 2))

	// 转化成小写字符
	var lower string
	lower = strings.ToLower(str)
	fmt.Println(lower)

	// 转化成大写字符
	var upper string
	upper = strings.ToUpper(str)
	fmt.Println(upper)

	// 剔除字符串开头和结尾的空白符号
	fmt.Println(strings.TrimSpace(str))

	// 分割字符串
	slice := strings.Fields(str)
	fmt.Println(slice)
	slice = strings.Split(str, ",")
	fmt.Println(slice)

	// 拼接字符串
	for _, val := range slice {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	fmt.Println(strings.Join(slice, ";"))
}
