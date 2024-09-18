package main

import (
	"fmt"
	"strings"
)

func main() {
	original := "Chan"
	cloned := strings.Clone(original)

	fmt.Println("Original:", original)
	fmt.Println("Cloned:", cloned)
	fmt.Println()

	fmt.Println("Checking if they have the same address or not.")
	if &original == &cloned {
		fmt.Println("Answer: Same address, brother")
	} else {
		fmt.Println("Answer: Same same but diffelant")
	}
	fmt.Println()

	cloned = "Myint"

	fmt.Println("Cloned value is changed.")
	fmt.Println("Original:", original)
	fmt.Println("Cloned:", cloned)
}
