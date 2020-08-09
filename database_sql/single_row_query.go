package main

import (
    "database/sql"
    "fmt"
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

    row := db.QueryRow("SELECT name FROM animals WHERE id=1")
    if err != nil {
        log.Fatal(err)
    }

    err = row.Scan(&s)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(s)
}