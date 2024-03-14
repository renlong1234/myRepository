package main

import "fmt"

func main() {
	k := 2
	nums := []int{1, 1, 1}
	//键盘输入数组（未完成 ）
	//var nums []int
	//var number int8
	//fmt.Println("请输入一个数组")
	//for i := 0; i < 3; i++ {
	//	fmt.Printf("nums[%v]=", i)
	//	fmt.Scanln(&number)
	//}
	fmt.Println(nums)
	ret := subarraySum(nums, k)
	fmt.Println(ret)
}
func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}
