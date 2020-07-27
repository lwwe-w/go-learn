package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

var name string = "ww12"

const pi = 3.1415

var (
	sex   string
	title string
)

const (
	n1 = iota //0
	n2        //1
	n3        //2
	n4        //3
)

func init() { // initialization of package

}

func main() {
	var goos string = runtime.GOOS
	var str string = "This is an example of a string"
	fmt.Println("operating", goos)
	path := os.Getenv("PATH")
	fmt.Println("path", path)
	fmt.Println(strings.HasPrefix(str, "Th"))
}
