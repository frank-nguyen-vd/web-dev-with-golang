package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("template.gohtml").ParseFiles("template.gohtml"))
}

func main() {
	/*
		List of predefined global functions: https://pkg.go.dev/text/template#hdr-Functions
	*/
	err := tpl.Execute(os.Stdout, []string{"hello", "world", "how are you"})
	if err != nil {
		log.Fatalln(err)
	}

}
