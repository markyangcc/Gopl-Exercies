// RUN with: "go run main.go"

package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer

	length := len(s)
	for i, ch := range s {
		if (length-i)%3 == 0 {
			buf.WriteRune(',')
		}
		buf.WriteRune(ch)
	}
	return buf.String()
}

func main() {
	s := comma("12345")
	fmt.Println(s)
}
