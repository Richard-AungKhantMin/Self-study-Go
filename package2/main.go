package main

import (
	"fmt"
	"selfstudy/package2/p1"
	"selfstudy/package2/p2"
)

func main() {

	var name string = "Richard"
	var age int = 22

	fmt.Println(p1.Func1(name))
	fmt.Println(p2.Func2(age))

}
