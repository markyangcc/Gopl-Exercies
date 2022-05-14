// print filename which have dup lines
//
// To end input from stdin, give EOF(end of file) to program
// EOF: press Ctrl + D(Linux), press Ctrl + C(Windows)
//
// run: go run main.go a.txt b.txt

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	files := os.Args[1:]

	// if files not provided as args, use stdin instead
	if len(files) == 0 {
		counts := make(map[string]int)
		countLines(os.Stdin, counts)
		for _, v := range counts {
			if v > 1 {
				fmt.Println("stdin")
				break
			}
		}

	} else {
		for _, arg := range files {
			counts := make(map[string]int)

			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)

			// loop map to check dup lines
			for _, v := range counts {
				if v > 1 {
					fmt.Println(arg)
					break
				}
			}
			f.Close()
		}
	}
}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
}
