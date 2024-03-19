package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	sum, area := 0, 0
	for left < right {
		if height[left] < height[right] {
			area = (right - left) * height[left]
			left++
		} else {
			area = (right - left) * height[right]
			right--
		}
		if sum < area {
			sum = area
		}
	}

	return sum
}

func main() {
	var height []int = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("最大装水的面积为：", maxArea(height))
}
