package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Start")

	content := "Chan Myint is one of the students from Gritlab"

	file, err := os.Create("./new_file.txt")
	isErrNill(err)

	length, err := io.WriteString(file, content)
	isErrNill(err)
	fmt.Println("The length of the content is:", length)

	each, err := os.ReadFile("./new_file.txt")
	isErrNill(err)
	fmt.Println(each)
}

func isErrNill(err error) {
	if err != nil {
		panic(err)
	}
}
