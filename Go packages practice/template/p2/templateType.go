package main

import (
	"fmt"
	"text/template/parse"
)

type common struct {
	description string
}

func main() {

	type Template struct {
		name string
		*parse.Tree
		*common
		leftDelim  string
		rightDelim string
	}

	// Assuming `parse.Tree` is already defined and instantiated elsewhere
	tree := &parse.Tree{} // Placeholder for example purposes
	commonData := &common{description: "A common description"}

	tmpl := Template{
		name:       "MyTemplate",
		Tree:       tree,
		common:     commonData,
		leftDelim:  "{{",
		rightDelim: "}}",
	}

	// Accessing fields and methods
	fmt.Println(tmpl.name)        // Output: MyTemplate
	fmt.Println(tmpl.description) // Accessing embedded `common` field
}
