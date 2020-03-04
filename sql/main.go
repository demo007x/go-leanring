package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db sql.DB

func main01() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_demo?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	// sql 预处理
	stmtlns, err := db.Prepare("insert into users values (?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtlns.Close()
	// Prepare statement for reading data
	stmtOut, err := db.Prepare("select * from users where name=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_demo?charset=utf8")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select name, age from users")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var name string
		var age int
		if err := rows.Scan(&name, &age); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("the name is %s, age is %d", name, age)
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
