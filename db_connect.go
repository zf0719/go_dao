package main

import (
	"database/sql"
	"log"

	_ "./go-sql-driver"
)

var db = &sql.DB{}

func Init() {
	var err error
	db, err = sql.Open("mysql", "root:Changeme_123@/mytestdb")
	if err != nil {
		db.Close()
		log.Fatal("Open database error: %s", err)
	}
}

func Fini() {
	db.Close()
}
