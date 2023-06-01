package main

import (
	"fmt"
	"github.com/printSANO/playWithGo/cal"
)

func main() {
	fmt.Println("Hello World")
	const a = 10
	const b = 11
	var c int
	c = cal.MySum(a, b)
	fmt.Println(c)
}