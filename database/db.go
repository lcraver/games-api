package database

import (
  "log"

  "database/sql"
  _ "github.com/lib/pq"
)

var (
    // DBCon is the connection handle
    // for the database
    DBCon *sql.DB
)

func CreateDB() {
    var err error
    DBCon, err = sql.Open("postgres","user=postgres password=abcd@1234 dbname=games_api sslmode=disable")

    err = DBCon.Ping()

    if err != nil {
      log.Fatal("Error: The data source arguments are not valid")
    }
}
