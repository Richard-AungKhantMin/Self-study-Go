package main

import (
	"fmt"
	"strings"
)

func main() {

	something := "Chan Myint is Chan Myint. Period! It's not Chan Chan nor Myint Myint."

	fmt.Println(something)
	fmt.Println()
	fmt.Println("Let's replace.")

	p1 := strings.Replace(something, "Chan Myint", "Jackie Chan", 1)
	p2 := strings.Replace(something, "Chan", "Ma", 2)
	p3 := strings.Replace(something, "Myint", "Nyein", -1)
	p4 := strings.Replace(something, "Chan", "Sayar Chan", -2)

	answers := []string{p1, p2, p3, p4}

	for n, each := range answers {
		fmt.Printf("Result %d : %s\n", n+1, each)
	}

}
