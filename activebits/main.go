package main

import (
	"fmt"
)

func main() {
	fmt.Println(ActiveBits(7))
}

func ActiveBits(n int) int {

	var count int

	for n > 0 {
		if n%2 == 1 {
			count++
		}
		n = n / 2
	}

	return count
}
