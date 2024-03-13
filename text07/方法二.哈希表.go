package main

import "fmt"

func main() {
	nums := []int{7, 11, 1, 2, 15, 5, 3}
	target := 9
	res := twoSum2(nums, target)
	fmt.Println(res)
}
func twoSum2(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
