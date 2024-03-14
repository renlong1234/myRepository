package main

import (
	"fmt"
	"sort"
	"strings"
)

// keyForString 生成字符串的键，用于字符异位分组
func keyForString(s string) string {
	// 将字符串转换为字符切片，并排序
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

// groupAnagrams 对给定的字符串切片进行字符异位分组
func groupAnagrams1(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, str := range strs {
		key := keyForString(str)
		groups[key] = append(groups[key], str)
	}

	var result [][]string
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groups := groupAnagrams1(strs)
	for _, group := range groups {
		fmt.Println(group)
	}
}
