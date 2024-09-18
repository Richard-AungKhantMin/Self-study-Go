package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("a", "z"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("z", "a"))
	// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
}
