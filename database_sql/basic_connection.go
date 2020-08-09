package main

import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "akalaj:@tcp(127.0.0.1:3306)/test")
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

    _, err = db.Exec("SELECT 'Hello World!'")
    if err != nil {
        log.Fatal(err)
    }
}