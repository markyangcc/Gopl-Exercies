// run with: go run main.go
// Input: "hello, 世界123"(press ctrl+D on linux to input EOF)

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	letters := make(map[rune]int)
	digits := make(map[rune]int)
	others := make(map[rune]int)

	invalid := 0

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			letters[r]++
		} else if unicode.IsDigit(r) {
			digits[r]++
		} else {
			others[r]++
		}

	}

	fmt.Printf("rune\tcount\n")
	fmt.Print("-----------letters--------------\n")
	for c, n := range letters {
		if c > 0 {
			fmt.Printf("%c\t%d\n", c, n)
		}
	}

	fmt.Print("-----------digits--------------\n")
	for c, n := range digits {
		if c > 0 {
			fmt.Printf("%c\t%d\n", c, n)
		}
	}

	fmt.Print("----------others--------------\n")
	for c, n := range others {
		if c > 0 {
			fmt.Printf("%c\t%d\n", c, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid utf8 characters\n", invalid)
	}

}
