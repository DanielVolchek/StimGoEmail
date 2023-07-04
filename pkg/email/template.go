package email

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

const HARD_CODED_PATH = "/Users/danielvolchek/code/stimstore/stim-email/templates/test.html"

func CreateTemplateTest(name, me string) {
	template, err := template.ParseFiles(HARD_CODED_PATH)

	if err != nil {
		log.Fatal(err)
	}

	result := new(bytes.Buffer)

	template.Execute(result, struct {
		Name string
		Me   string
	}{name, me})
	fmt.Println(result.String())
}
