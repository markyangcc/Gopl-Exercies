// run with "go run *.go" or just "go run ."

package main

import "fmt"

func main() {
	// 120(10) = 1111000(2)
	var x uint64 = 120

	if PopCount(x) != 4 {
		fmt.Println("failed")
	} else {
		fmt.Println("pass")
	}
}
