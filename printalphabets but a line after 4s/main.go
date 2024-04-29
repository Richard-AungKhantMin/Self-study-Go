package main

import "github.com/01-edu/z01"

func main() {

	var count int = 0

	for i := 'a'; i <= 'z'; i++ {
		count++
		z01.PrintRune(i)
		if count%4 == 0 {
			z01.PrintRune('\n')
		}
	}

}
