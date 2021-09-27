package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var fm = template.FuncMap{
	"mul2": func(x float64) float64 { return x * 2 },
	"div2": func(x float64) float64 { return x / 2 },
	"sqrt": math.Sqrt,
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("template.gohtml").Funcs(fm).ParseFiles("template.gohtml"))
}

func main() {

	err := tpl.Execute(os.Stdout, 32.0)
	if err != nil {
		log.Fatalln(err)
	}

}
