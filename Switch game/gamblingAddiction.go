package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("5€ for each ticket!!!")
	fmt.Println()
	fmt.Print("Your number: ")

	r := rand.Intn(100)

	switch {
	default:
		fmt.Println(r)
		fmt.Println("Maybe next time, my love.")

	case r%3 == 0:
		fmt.Println(r)
		fmt.Println("You won 3€.")

	case r%5 == 0:
		fmt.Println(r)
		fmt.Println("You won 5€.")

	case r%37 == 0:
		fmt.Println(r)
		fmt.Println("You won 50€.")

	case r%73 == 0:
		fmt.Println(r)
		fmt.Println("You won 100€.")

	}
}
