package main

import "fmt"

func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func main() {
	var n int
	fmt.Printf("Please enter the number of stairs n=：")
	fmt.Scanln(&n)
	fmt.Printf("一共有%v种方式爬楼梯", climbStairs(n))
}
