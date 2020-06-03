package main

import (
	"html/template"
	"os"
)

type name struct {
	Name string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := name{
		Name: "<script>alert('Howdy!');</script>",
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}