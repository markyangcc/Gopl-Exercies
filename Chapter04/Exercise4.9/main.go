package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	letters := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		letters[word]++
	}

	for w, n := range letters {
		fmt.Printf("%s %d\n", w, n)
	}

}
