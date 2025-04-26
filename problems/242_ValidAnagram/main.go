package main

// Array would be faster to lookup than map
// due to the constraint of the problem, we can array of alphabet
// time  O(n)
// space O(n)
//
// remember to create array when dealing with string
func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	arr := [26]int{}
	for _, o := range s {
		arr[o-'a']++
	}
	for _, o := range t {
		arr[o-'a']--
		if arr[o-'a'] < 0 {
			return false
		}
	}
	return true
}

func main() {
	isAnagram("anagram", "nagaram")
}
