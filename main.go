package main

import (
	"fmt"
	"strings"
	"net/http"
	"github.com/gorilla/sessions"
	"database/sql"
)

var store = sessions.NewCookieStore([]byte("app"))
var config instanceConfig
var db *sql.DB

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hi there, I love %s!", r.URL.Path[1:])

	if strings.Contains(r.URL.Path[1:], "example") {
		example_controller(w, r)
	}
}

func main() {
	config.readConfiguration();

	var err error
	db, err = sql.Open(config.Driver, config.Username + ":@tcp(" + config.Hostname + ":" + config.Port + ")/" + config.Database)

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":" + config.ServerPort, nil)
}
