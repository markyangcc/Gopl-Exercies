package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}

}
