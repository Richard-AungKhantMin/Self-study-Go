package main

import (
	"fmt"
)

func main() {

	fmt.Println(base2(7))
	fmt.Println(base2(13))
	fmt.Println(base2(32))
	fmt.Println(base2(56))
	fmt.Println(base2(1024))
	fmt.Println(base2(1717))
	fmt.Println(base2(19))
	fmt.Println(base2(71))
}

func base2(n int) string {

	var slicy []int
	var answer string

	for n > 0 {
		slicy = append(slicy, n%2)
		n = n / 2
	}

	for i := len(slicy) - 1; i >= 0; i-- {
		answer = answer + string(rune(slicy[i]+'0'))
	}

	return answer
}
