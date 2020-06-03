package main

import (
	"html/template"
	"os"
)

type name struct {
	Name    string
	Program string
	Number int
	Float float64
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := name{
		Name:    "John Smith",
		Program: "Golang",
		Number: 1027,
		Float: 10.27,
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
