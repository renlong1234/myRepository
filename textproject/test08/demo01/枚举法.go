package main

import "fmt"

func main() {
	k := 2
	nums := []int{1, 1, 1}
	ret := subarraySum(nums, k)
	fmt.Println(ret)
}
func subarraySum(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j >= 0; j-- {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}
