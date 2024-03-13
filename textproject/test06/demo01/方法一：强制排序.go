package main

import "fmt"

func main() {
	nums := []int{1, 2, 4, 6, 7}
	target := 3
	ret := TwoSum(nums, target)
	fmt.Println(ret)
}
func TwoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
