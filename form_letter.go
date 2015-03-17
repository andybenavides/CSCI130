package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Define the template
	const letter = `
Dear {{.Name}},
{{if .Attended}}
Thanks for stopping by.{{else}}
Maybe next time?{{end}}
{{with .Beverage}}Thank you for the tasty{{.}}.
{{end}}
Best wishes,
Jose
`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Beverage string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Justin", "Vodka", true},
		{"Anthony", "Pineapple Juice", true},
		{"Tyler", "", true},
		{"Carlos", "Tequila", true},
	}

	// STEP 1 & STEP 2
	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	//STEP 3
	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}
