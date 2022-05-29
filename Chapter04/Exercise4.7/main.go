package main

import (
	"fmt"
	"unicode/utf8"
)

// classic method: mutil-reverse to make this
func reverseUTF8Str(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverse(b[i : i+size])
		i += size
	}
	reverse(b)

	return b
}
func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func main() {
	b := []byte("hello, 世界")
	reverseUTF8Str(b)

	fmt.Printf("%s\n", b)
}
