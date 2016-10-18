package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"encoding/json"
)

type exampleModel struct {
	Id int
    Name string
    Surname string
}

func (m *exampleModel) setName(Name string) {
	m.Name = Name
}

func (m exampleModel) getName() string {
	return m.Name
}

func (m *exampleModel) setSurname(Surname string) {
	m.Surname = Surname
}

func (m exampleModel) getSurname() string {
	return m.Surname
}

func (m exampleModel) selectFrom(w http.ResponseWriter) {
	var model exampleModel

	mlist := []exampleModel{}

	rows, err := db.Query("SELECT * FROM example")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
        err = rows.Scan(&model.Id, &model.Name, &model.Surname)
        if err != nil {
			panic(err)
		}

		mlist = append(mlist, model)
    }

    b, err := json.Marshal(mlist)
    fmt.Fprintln(w, string(b))
}

func (m exampleModel) insert() {
	stmt, err := db.Prepare("INSERT example SET name=?, surname=?")

	if err != nil {
		panic(err)
	}

    res, err := stmt.Exec(m.Name, m.Surname)

    if err != nil {
		panic(err)
	}

    id, err := res.LastInsertId()

    if err != nil {
		panic(err)
	}

    fmt.Println(id)
}

func (m exampleModel) update() {
	stmt, err := db.Prepare("UPDATE example SET name=?,surname=? WHERE id=?")

	if err != nil {
		panic(err)
	}

    res, err := stmt.Exec(m.Name, m.Surname, m.Id)

    if err != nil {
		panic(err)
	}

    id, err := res.LastInsertId()

    if err != nil {
		panic(err)
	}

    fmt.Println(id)
}

func (m exampleModel) delete(id int) {
	stmt, err := db.Prepare("DELETE FROM example WHERE id=?")

	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(id)

    if err != nil || res != nil {
		panic(err)
	}

	// statusCode := "202"

	// m := Message{statusCode}
	// b, err := json.Marshal(m)

	// if err != nil {
 //        fmt.Println(err)
 //        return
 //    }
}
