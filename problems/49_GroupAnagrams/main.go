package main

import (
	"fmt"
	"leetcode/lib/sort"
)

// place the sorted string to hash map
func groupAnagrams(strs []string) [][]string {

	hash := make(map[string][]string)
	for _, str := range strs {
		sorted := sort.SortString(str)
		hash[sorted] = append(hash[sorted], str)
	}
	ans := [][]string{}
	for _, v := range hash {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	result := groupAnagrams([]string{"bat", "tab", "mad"})
	fmt.Println(result)
}
