package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
    "database/sql"
)

type exampleModel struct {
    Name string
    Surname string
}

func (m exampleModel) setName() string (name string) {
	return m.name;
}

func getName (model exampleModel) {
	return model.Name;
}

func getSurname (model exampleModel) {
	return model.Surname;
}

func create (model exampleModel) {

}

func insert (model exampleModel) {

}

func update (model exampleModel) {
	//fmt.println("Printing model")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/example")

	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")

	if err != nil {
		panic(err)
	}

    res, err := stmt.Exec(model.Name, model.Surname, "")

    if err != nil {
		panic(err)
	}

    id, err := res.LastInsertId()

    if err != nil {
		panic(err)
	}

    fmt.Println(id)

}

func delete (model exampleModel) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/example")

	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("DELETE FROM userinfo WHERE id=?")

	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(model.Name, model.Surname, "")

    if err != nil {
		panic(err)
	}
}