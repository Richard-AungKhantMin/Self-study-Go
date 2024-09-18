package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		log.Println("Please put only 2 inputs.")
		return
	}

	log.Println(strings.Contains(os.Args[1], os.Args[2]))
}
