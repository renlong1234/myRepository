package main

import "fmt"

func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
func main() {
	nums := []int{0, 1, 0, 3, 12}
	fmt.Println("原始数组为:", nums)
	moveZeroes(nums)
	fmt.Println("移动之后的数组为：", nums)
}
