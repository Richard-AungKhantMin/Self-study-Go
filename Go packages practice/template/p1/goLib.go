package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Define a template.
	const letter = `
Dear {{.Name}}
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}

	var boolVar bool
	switch os.Args[3] {
	case "true":
		boolVar = true
	case "false":
		boolVar = false
	}
	name := os.Args[1]
	gift := os.Args[2]
	var recipients = []Recipient{

		{name, gift, boolVar},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}
