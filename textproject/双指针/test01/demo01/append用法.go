package main

import "fmt"

func moveZeroes(nums []int) {
	var nums2 []int
	j := 0
	for _, num := range nums {
		if num != 0 {
			nums2 = append(nums2, num)
			j++
		}
	}
	for k := j; k < len(nums); k++ {
		nums2 = append(nums2, 0)
	}
	copy(nums, nums2)

}
func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
