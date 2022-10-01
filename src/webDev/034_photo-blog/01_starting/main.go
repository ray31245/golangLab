package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
)

var db *sql.DB
var err error

func index(w http.ResponseWriter, req http.Request) {
	_, err := io.WriteString(w, "at index")
	check(err)
}

func amigos(w http.Response, req *http.Request) {
	rows, err := db.Query(`SELECT aName FROM amigos;`)
	check(err)
	defer rows.Close()
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
