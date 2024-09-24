//os.Create
//os.ReadFile
//io.WriteString

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("1. The program starts here.")
	//prints here

	content := "Chan Myint is one of the students from Gritlab"

	file, err := os.Create("./new_file.txt")
	isErrNill(err)
	fmt.Println("2. ", file)
	//prints here

	length, err := io.WriteString(file, content)
	isErrNill(err)
	fmt.Println("3. The length of the content is:", length)
	//prints here

	each, err := os.ReadFile("./new_file.txt")
	isErrNill(err)
	fmt.Println("4. ", string(each))
	// prints here

	new, err := os.Create("./new_file.txt")
	isErrNill(err)
	fmt.Println("5. ", new)
	//prints here

	new1, err := os.ReadFile("./new_file.txt")
	isErrNill(err)
	fmt.Println("6. ", new1)
	//prints here

	fmt.Println("7. ", string(each))
	// prints here
}

func isErrNill(err error) {
	if err != nil {
		panic(err)
	}
}
