package main

import (
	"log"
	"os"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {
	buddha := []sage{
		{Name: "Buddha",
			Motto: "The belief of no beliefs"},
		{Name: "Jesus",
			Motto: "I love everyone"},
		{Name: "LKY",
			Motto: "No corruption"},
	}

	err := tpl.Execute(os.Stdout, buddha)
	if err != nil {
		log.Fatalln(err)
	}

}
