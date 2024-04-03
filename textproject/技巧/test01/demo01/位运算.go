package main

import "fmt"

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

func main() {
	var n int
	fmt.Printf("请输入n=")
	fmt.Scanln(&n)
	//var nums []int{}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("请输入数组nums[%d]=", i)
		fmt.Scanln(&nums[i])
	}

	fmt.Println(singleNumber(nums))
}
