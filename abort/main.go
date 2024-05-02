package main

import (
	"fmt"
)

func main() {
	middle := Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}

func Abort(a, b, c, d, e int) int {

	var slicy []int

	slicy = append(slicy, a)
	slicy = append(slicy, b)
	slicy = append(slicy, c)
	slicy = append(slicy, d)
	slicy = append(slicy, e)

	for i := 0; i < len(slicy)-1; i++ {
		for j := i + 1; j < len(slicy); j++ {
			if slicy[i] > slicy[j] {
				slicy[i], slicy[j] = slicy[j], slicy[i]
			}
		}
	}

	return slicy[2]
}
