package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Message struct {
    Name string
    Body string
}

func example_controller(w http.ResponseWriter, r *http.Request) {
	example := exampleModel { Name: "Alejandro", Surname: "Sanchez" }

	fmt.Fprintf(w, example.Name)
	fmt.Fprintf(w, example.Surname)

	write(example)

	m := Message{"Status", "400"}

	b, err := json.Marshal(m)

	fmt.Fprint(w, b)

	if err != nil {
		panic(err)
	}
}
