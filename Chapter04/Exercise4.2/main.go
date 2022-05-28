package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func hashSHAxxx(bits string, stdin string) {

	if bits == "384" {
		c := sha512.Sum384([]byte(stdin))
		fmt.Printf("%x\n", c)
	} else if bits == "512" {
		c := sha512.Sum512([]byte(stdin))
		fmt.Printf("%x\n", c)
	} else {
		c := sha256.Sum256([]byte(stdin))
		fmt.Printf("%x\n", c)
	}
}

func parseSHAType() string {
	var shatype string
	for i, v := range os.Args {
		if v == "-t" {
			shatype = os.Args[i+1]
		}
	}
	return shatype
}

func main() {
	// read from stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter text: ")
	text, _ := reader.ReadString('\n')

	// detect sha type from cmd args
	shatype := parseSHAType()
	if shatype != "256" && shatype != "384" && shatype != "512" {
		fmt.Printf("SHA%s not support..., SHA256 used instead\n", shatype)
		shatype = "256"
		fmt.Printf("SHA%s used...\n", shatype)
	} else {
		fmt.Printf("SHA%s detected and used...\n", shatype)
	}

	// hash stdin
	hashSHAxxx(shatype, text)
}
