package main

import (
	"html/template"
	"log"
	"os"
)

type hotel struct {
	Name, Address, City, Zip, Region string
}

type hotels []hotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := hotels{
		hotel{
			Name:    "Kazakzhstan",
			Address: "Kazibek bi,23",
			City:    "Almaty",
			Zip:     "01233923",
			Region:  "02",
		},
		hotel{
			Name:    "California",
			Address: "MayamiBeach,23",
			City:    "California",
			Zip:     "1234123413",
			Region:  "USA",
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
