package main

import (
	"fmt"
	"strings"
	"net/http"
)

type Message struct {
    Status string
}

var example exampleModel

func example_controller(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(example.Name)
	fmt.Printf(example.Surname)

	if strings.Contains(r.URL.Path[1:], "select") {
		fmt.Printf("Select")
		example.selectFrom(w)
	}

	if strings.Contains(r.URL.Path[1:], "insert") {
		fmt.Printf("Insert")
		example := exampleModel { Name: "Alejandro", Surname: "Sanchez" }
		example.insert()
	}

	if strings.Contains(r.URL.Path[1:], "update") {
		fmt.Printf("Update")
		example := exampleModel { Id: 13, Name: "Carlos", Surname: "Sanchez" }
		example.update()
	}

	if strings.Contains(r.URL.Path[1:], "delete") {
		fmt.Printf("Delete")
		example.delete(1)
	}
}
