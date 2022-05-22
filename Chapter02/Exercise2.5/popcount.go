package main

func PopCount(num uint64) (ones int) {
	for ; num > 0; num &= num - 1 {
		ones++
	}
	return
}
