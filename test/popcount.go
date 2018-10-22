package main

import "fmt"

var pc [256]byte

func init() {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	count := 0
	for i:=uint(0); i<8; i++ {
		count += int(pc[byte(x >> (i*8))])
	}
	return count
}

func main() {
	n := uint64(5)
	count := PopCount(n)
	fmt.Println("the time" ,count)
}


