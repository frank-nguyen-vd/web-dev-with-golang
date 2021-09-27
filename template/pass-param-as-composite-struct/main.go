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

type car struct {
	Brand string
	Doors int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {
	sages := []sage{
		{Name: "Buddha",
			Motto: "The belief of no beliefs"},
		{Name: "Jesus",
			Motto: "I love everyone"},
		{Name: "LKY",
			Motto: "No corruption"},
	}

	cars := []car{
		{Brand: "BWM", Doors: 4},
		{Brand: "Ferrari", Doors: 2},
	}

	data := items{
		Wisdom: sages, Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

}
