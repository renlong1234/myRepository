package main

import "fmt"

func sum(num1 int, num2 int) int {
	s := num1 + num2
	return s
}
func main() {
	s := sum(10, 20)
	fmt.Printf("sum=%v\n", s)
	num1, _ := strconv.Atoi("765")
	fmt.Println(num1)
	str1 := strconv.Itoa(89)
	fmt.Println(str1)
	fmt.Printf("num1的数据类型是：%T,str1的数据类型是：%T", num1, str1)

	flag := strings.Contains("sadasjgojkgokj", "go")
	count := strings.Count("sadasjgojkgokj", "go")
	fmt.Println()
	fmt.Println(count)
	fmt.Printf("%v", flag)

	fmt.Println()
	fmt.Println(strings.Index("lagond", "go"))
}
