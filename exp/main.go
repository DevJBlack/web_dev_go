package main

import (
	"html/template"
	"os"
)

type dog struct {
	Name string
}

type user struct {
	Dog     dog
	Name    string
	Program string
	Number  int
	Float   float64
	Slice   []string
	Map     map[string]string
}

func main() {
	// Takes in as many files as we want to pass in and returns a pointer to that file and and error
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := user{
		Dog: dog{
			Name: "Levi",
		},
		Name:    "John Smith",
		Program: "Golang",
		Number:  1027,
		Float:   10.27,
		Slice:   []string{"Name", "City", "Town"},
		Map: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
	}

	// os.Stdout is to the terminal // t is the template that was made
	// First argument is where we want to write the output which is os.Stdout, second argument is the data we pass in
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	// If we copy the execute function and give a data struct a new value, it will print it twice, one with John Smith, the other with Devin Black
	// data.Name = "Devin Black"
	// err = t.Execute(os.Stdout, data)
	// if err != nil {
	// 	panic(err)
	// }

	// values := map[string]int{
	// 	"Golang":     100,
	// 	"JavaScript": 90,
	// 	"Python":     50,
	// }

	// for k, v := range values {
	// 	fmt.Printf("Key: %v, Value: %v", k, v)
	// }

}
