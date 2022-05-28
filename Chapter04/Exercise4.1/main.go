package main

import (
	"crypto/sha256"
	"fmt"
)

// bitDiff for any []byte
func bitDiff(a, b []byte) int {
	var cnt int
	for i := 0; i < len(a); i++ {
		res := a[i] ^ b[i]
		cnt += popCount(res)
	}
	return cnt
}

// count numbers of bit 1 in one byte(aka uint8)
func popCount(num uint8) int {
	var cnt int
	for num != 0 {
		num = num & (num - 1)
		cnt++
	}
	return cnt
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("y"))

	// %x for hex format
	fmt.Printf("c1: %x\n", c1)
	fmt.Printf("c2: %x\n", c2)

	// Check if the function is correct or not, short bits for human
	c3 := []byte{1, 0, 0, 1}
	c4 := []byte{1, 0, 1, 1}

	// %b for bit format
	fmt.Printf("c3: %b\n", c3)
	fmt.Printf("c4: %b\n", c4)

	cntsha256 := bitDiff(c1[:], c2[:])
	cntshort := bitDiff(c3, c4)

	fmt.Println(cntsha256)
	fmt.Println(cntshort)

}
