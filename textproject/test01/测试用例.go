package main

import "fmt"

func sum(num1 int, num2 int) int {
	s := num1 + num2
	return s
}
func main() {
	s := sum(10, 20)
	fmt.Printf("sum=%v\n", s)
}
