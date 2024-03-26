package main

import "fmt"

func bubbleSort(arr []int, n int) []int {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
func main() {
	arr := []int{2, 9, 3, 4, 5, 1, 7, 8, 0}
	k := len(arr)
	fmt.Println(bubbleSort(arr, k))
}
