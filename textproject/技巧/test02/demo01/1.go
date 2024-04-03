package main

import "fmt"

func majorityElement(nums []int) (ans int) {
	cnts := 0
	for _, v := range nums {
		if v == ans {
			cnts++
		} else if cnts == 0 {
			ans = v
		} else {
			cnts--
		}
	}
	return
}

func main() {
	var n int
	fmt.Printf("请输入n=%d", n)
	fmt.Scanln(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("nums[%d]", i)
		fmt.Scanln(&nums[i])
	}
	fmt.Println(majorityElement(nums))
}
