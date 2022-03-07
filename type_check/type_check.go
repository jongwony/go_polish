package main

import "fmt"

func main() {
	var x interface{}
	x = 4
	s, ok := x.(int)
	fmt.Println(4 == 4.0)
	fmt.Println(s, x, ok)
}
