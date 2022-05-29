package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func replaceUnicodeSpace(b []byte) []byte {

	// traversal one by one, not efficient but easy to understand
	for i := 0; i < len(b); {
		first, size := utf8.DecodeRune(b[i:])
		if unicode.IsSpace(first) {
			next, _ := utf8.DecodeRune(b[i+size:])
			if unicode.IsSpace(next) {
				copy(b[i:], b[i+size:])
				b = b[:len(b)-size] // update slice b
				continue
			}
		}
		i += size
	}

	return b
}

func main() {
	b := []byte("  hello,  世界   ")
	b = replaceUnicodeSpace(b)
	fmt.Printf("%s", b)
}
