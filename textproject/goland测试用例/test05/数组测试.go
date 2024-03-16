package main

import "fmt"

func main() {
	var arr [3]int8
	fmt.Println("数组arr的地址是:%p", &arr[0])
	fmt.Println("数组arr的地址是:%p", &arr[1])
	fmt.Println("数组arr的地址是:%p", &arr[2])
	var arr1 = [3]int{3, 8, 9}
	test1(arr1)
	fmt.Println(arr1)
	test2(&arr1)
	fmt.Println(arr1)
}
func test1(arr [3]int) {
	arr[0] = 7
}
func test2(arr *[3]int) {
	(*arr)[0] = 7
}
