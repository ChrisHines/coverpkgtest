package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl = template.Must(template.New("tmpl.txt").ParseFiles("tmpl.txt"))

func main() {
	if err := tmpl.Execute(os.Stdout, 9); err != nil {
		log.Fatal(err)
	}
}
