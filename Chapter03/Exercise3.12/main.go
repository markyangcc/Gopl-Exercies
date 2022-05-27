package main

import (
	"fmt"
)

func isAnagram(s1, s2 string) bool {
	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	for _, ch := range s1 {
		map1[ch]++
	}
	for _, ch := range s2 {
		map2[ch]++
	}

	for k, v := range map1 {
		if v != map2[k] {
			return false
		}
	}

	// if s1 == s2, return false
	return s1 != s2
}

func main() {
	fmt.Println(isAnagram("aba", "baa"))
	fmt.Println(isAnagram("aaa", "aaa"))
	fmt.Println(isAnagram("aaa", "aab"))
}
