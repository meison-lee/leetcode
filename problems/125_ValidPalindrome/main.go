package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	runes := []rune(strings.ToLower(s))
	if len(runes) < 2 {
		return true
	}
	left, right := 0, len(runes)-1

	for left < right {
		if runes[left] < 'a' || runes[left] > 'z' {
			left++
			continue
		}
		if runes[right] < 'a' || runes[right] > 'z' {
			right--
			continue
		}
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func main() {
	result := isPalindrome("0P")
	fmt.Println(result)
}
