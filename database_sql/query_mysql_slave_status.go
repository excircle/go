package main

import (
  "database/sql"
  "log"
  _ "github.com/go-sql-driver/mysql"
  "fmt"
  "strings"
)

func main() {

  //Open a connection to MySQL
  db, err := sql.Open("mysql", "username:password@tcp(hostname:mysql_port)/db_name")
  if err != nil {
    log.Fatal(err)
  }

  //No matter what happens, execute a db.Close() as the last thing we do
  defer db.Close()


  //Build sql.Rows object which contains our query data
  rows, err := db.Query("SHOW SLAVE STATUS")

  if err != nil {
    log.Fatal(err)
  }

  // sql.Rows has a function which returns all column names
  // as a slice of []string. Variable "columns" represents this
  columns, err := rows.Columns()
  if err != nil {
    log.Fatal(err)
  }

  // variable "values" is a pre-populated array of empty interfaces
  // We load an empty interface for every column 'sql.Rows' has.
  // The interfaces will allow us to call methods of any type that replaces it
  values := make([]interface{}, len(columns))

  // for every key we find while traversing array "columns"
  // set the corresponding interface in array "values" to be populated
  // with an empty sql.RawBytes type
  // sql.RawBytes is analogous to []byte
  for key, _ := range columns {
    values[key] = new(sql.RawBytes)
  }

  //Contrary to appearances, this is not a loop through every row
  // "rows.Next()"" is a recursive function that is called immediately
  // upon every row until we hit "rows.Next == false"
  // This is important because it means you must prepopulate variables or
  // arrays to the exact number of columns in the target SQL table
  // more details at: https://golang.org/pkg/database/sql/#Rows.Next
  for rows.Next() {
    //the "values..." tells Go to use every available slot to populate data
    err = rows.Scan(values...)
    if err != nil {
      log.Fatal(err)
    }
  }

  for index, columnName := range columns {
    // convert sql.RawBytes to String using "fmt.Sprintf"
    columnValue := fmt.Sprintf("%s", values[index])
    
    // Remove "&" from row values
    columnValue = strings.Replace(columnValue, "&", "", -1)
    
    // Optional: Don't display values that are NULL
    // Remove "if" to return empty NULL values
    if len(columnValue) == 0 {
      continue
    }

    // Print finished product
    fmt.Printf("%s: %s\n", columnName, columnValue)
  }
}
