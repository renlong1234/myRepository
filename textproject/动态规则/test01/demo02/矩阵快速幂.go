package main

import "fmt"

type matrix [2][2]int

func mul(a, b matrix) (c matrix) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][0]*b[0][j] + a[i][1]*b[1][j]
		}
	}
	return c
}

func pow(a matrix, n int) matrix {
	res := matrix{{1, 0}, {0, 1}}
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
	}
	return res
}

func climbStairs(n int) int {
	res := pow(matrix{{1, 1}, {1, 0}}, n)
	return res[0][0]
}

func main() {
	var n int
	fmt.Printf("Please enter the number of stairs n=：")
	fmt.Scanln(&n)
	fmt.Printf("一共有%v种方式爬楼梯", climbStairs(n))
}
