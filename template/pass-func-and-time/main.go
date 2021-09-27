package main

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var fm = template.FuncMap{
	"uc":       strings.ToUpper,
	"ft":       getAlias,
	"fdateMDY": formatMDY,
}

func getAlias(name string) string {
	name = strings.TrimSpace(name)
	return name[:3]
}

func formatMDY(t time.Time) string {
	// Ref: 01/02 03:04:05PM '06 -0700
	// Where
	// 01 : month
	// 02 : day
	// 03 : hour
	// 04 : minute
	// 05 : second
	// 06 : year
	// 07 : timezone

	return t.Format("01-02-2006 07")
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("template.gohtml").Funcs(fm).ParseFiles("template.gohtml"))
}

func main() {

	err := tpl.Execute(os.Stdout, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

}
