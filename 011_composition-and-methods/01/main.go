package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p1 := person{
		Name: "James Bond",
		Age:  42,
	}

	p2 := person{
		Name: "Akzhol Kanybekuly",
		Age:  28,
	}

	arr := []person{p1, p2}

	err := tpl.Execute(os.Stdout, arr)
	if err != nil {
		log.Fatalln(err)
	}

}
