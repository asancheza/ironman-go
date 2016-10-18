package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type Message struct {
    Status string
}

func example_controller(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	
    if err != nil {
        panic(err)
    }

    fmt.Fprintln(w, string(body))

	if strings.Contains(r.URL.Path[1:], "select") {
		fmt.Printf("Select")
		var selectModel exampleModel;

		selectModel.selectFrom(w)
	}

	if strings.Contains(r.URL.Path[1:], "insert") {
		fmt.Printf("Insert")
		var insertModel exampleModel;
		err = json.Unmarshal([]byte(body), &insertModel)
		fmt.Fprintln(w, insertModel)

		insertModel.insert()
	}

	if strings.Contains(r.URL.Path[1:], "update") {
		fmt.Printf("Update")
		var updateModel exampleModel;
		err = json.Unmarshal([]byte(body), &updateModel)
		fmt.Fprintln(w, updateModel)
		updateModel.update()
	}

	if strings.Contains(r.URL.Path[1:], "delete") {
		fmt.Printf("Delete")
		var deleteModel exampleModel;
		err = json.Unmarshal([]byte(body), &deleteModel)
		fmt.Fprintln(w, deleteModel)

		deleteModel.delete(deleteModel.Id)
	}
}
