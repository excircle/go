package main

import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

var s string

func main() {
    db, err := sql.Open("mysql", "akalaj:@unix(/var/run/mysqld/mysqld.sock)/test")
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

    err = db.Ping()
    if err == nil {
        log.Println("Database connection established!")
    } else {
        log.Println(err)
    }
}