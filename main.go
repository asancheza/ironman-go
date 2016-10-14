
package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"net/http"
	"gopkg.in/yaml.v2"
)

type instanceConfig struct {
    Hostname string
    Username string
    Password string
    Driver   string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	if r.URL.Path[1:] == "example" {
		example_controller(w, r)
	}
}

func main() {
	configFile, _ := filepath.Abs("./config.yml")
	yamlFile, err := ioutil.ReadFile(configFile)

	if err != nil {
		panic(err)
	}

	var config instanceConfig

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Value:  %#v\n", config)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
