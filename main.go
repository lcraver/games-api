package main

import (
  "net/http"
  "log"

  "database/sql"
  _ "github.com/lib/pq"
  "fmt"

  "github.com/lcraver/games-api/models"
)

var db *sql.DB

func main() {
  router := NewRouter()
  createDB()
  log.Fatal(http.ListenAndServe(":80", router))
}

func createDB() {
    var err error
    db, err = sql.Open("postgres","user=postgres password=abcd@1234 dbname=games_api sslmode=disable")

    err = db.Ping()

    if err != nil {
      log.Fatal("Error: The data source arguments are not valid")
    }

    GetAll("games")
}

func GetAll(table string) {
  results, err := db.Query("SELECT * FROM "+table)

  games := make([]*models.Game, 0)
  for results.Next() {
    game := new(models.Game)
    err := results.Scan(&game.Name, &game.Id)
    if err != nil {
      fmt.Println(err);
      return
    }
    games = append(games, game)
  }

  for i := 0; i < len(games); i++ {
		fmt.Println(games[i]);
	}

  if err != nil {
    log.Fatal(err)
  }
}
