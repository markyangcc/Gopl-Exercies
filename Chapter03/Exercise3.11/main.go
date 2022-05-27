// RUN with: "go run main.go"

package main

import (
	"bytes"
	"fmt"
	"strings"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func comma(s string) string {
	var buf bytes.Buffer

	// handle symbol +/- first
	if string(s[0]) == "+" || string(s[0]) == "-" {
		s = s[1:]
	}

	// WriteRune() stop in front of dot in float num
	index := strings.Index(s, ".")
	length := index

	// if strings.Index() not found, means the dot located at the tail
	if index == -1 {
		length = len(s)
	}

	for i, ch := range s {
		if i < length {
			if (length-i)%3 == 0 {
				buf.WriteRune(',')
			}
		}
		buf.WriteRune(ch)
	}

	return buf.String()
}

func main() {
	s := comma("12345678.987654")
	s1 := comma("+12345678.987654")
	s2 := comma("12345678")
	s3 := comma("+12345678")

	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
