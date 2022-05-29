package main

import "fmt"

func rotate(s *[]int, pos int) {
	tmp := make([]int, len(*s))
	copy(tmp[:], (*s)[pos:])
	copy(tmp[(len(*s)-pos):], (*s)[:pos])
	// swap two array ptr
	*s = tmp
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	rotate(&a, 2)
	fmt.Println(a)
}
