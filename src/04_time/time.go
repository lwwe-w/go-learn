package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())

	t = time.Now().UTC()
	fmt.Println(t)

	week = 60 * 60 * 24 * 7 * 1e9
	weekFromNow := t.Add(time.Duration(week))
	fmt.Println(weekFromNow)

	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
}
