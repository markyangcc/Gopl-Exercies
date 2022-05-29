package main

import "fmt"

func removeDupStr(strs []string) []string {

	nums := 0
	for i, j := 0, 0; j < len((strs)); {

		if i != j && strs[i] != strs[j] {
			strs[i+1] = strs[j]
			i++
			nums++
		}

		if strs[i] == strs[j] {
			j++
		}
	}

	return strs[:nums]
}

func main() {

	strs := []string{"aa", "aa", "bb", "cc", "aa"}
	fmt.Println(strs)

	strs = removeDupStr(strs)
	fmt.Println(strs)
}
