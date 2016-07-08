package handlers

import (
  "encoding/json"
  "net/http"

  "github.com/gorilla/mux"

  _ "github.com/lib/pq"
  "fmt"
  "log"

  "github.com/lcraver/games-api/models"
  "github.com/lcraver/games-api/database"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
  results, err := database.DBCon.Query("SELECT * FROM games")

  games := make([]*models.Game, 0)
  for results.Next() {
    game := new(models.Game)
    err := results.Scan(&game.Id, &game.Title)
    if err != nil {
      fmt.Println(err);
      return
    }
    games = append(games, game)
  }

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(games); err != nil {
      panic(err)
  }

  if err != nil {
    log.Fatal(err)
  }
}

func GetByName(w http.ResponseWriter, r *http.Request) {

  vars := mux.Vars(r)
  var title = vars["title"]

  game := new(models.Game);

  query, err := database.DBCon.Prepare("SELECT * FROM games WHERE title=$1")
  err = query.QueryRow(title).Scan(&game.Id, &game.Title);

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(game); err != nil {
      panic(err)
  }

  if err != nil {
    log.Fatal(err)
  }
}
