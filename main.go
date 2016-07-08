package main

import (
  "net/http"
  "log"

  "github.com/lcraver/games-api/database"
)

func main() {
  router := NewRouter()
  database.CreateDB()
  log.Fatal(http.ListenAndServe(":80", router))
}
