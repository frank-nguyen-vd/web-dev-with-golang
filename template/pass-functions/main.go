package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": getAlias,
}

func getAlias(name string) string {
	name = strings.TrimSpace(name)
	return name[:3]
}

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
	tpl = template.Must(template.New("template.gohtml").Funcs(fm).ParseFiles("template.gohtml"))
}

func main() {
	sages := []sage{
		{Name: "Buddha",
			Motto: "The belief of no beliefs"},
		{Name: "Jesus Christ",
			Motto: "I love everyone"},
		{Name: "Lee Kuan Yew",
			Motto: "No corruption"},
	}

	cars := []car{
		{Brand: "BWM", Doors: 4},
		{Brand: "FERRARI", Doors: 2},
		{Brand: "MERCEDES", Doors: 3},
	}

	data := items{
		Wisdom: sages, Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

}
