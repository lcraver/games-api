package handlers

import (
  "encoding/json"
  "net/http"

  //"github.com/gorilla/mux"

  "github.com/lcraver/games-api/models"
)


func GameHandler(w http.ResponseWriter, r *http.Request) {

  //vars := mux.Vars(r)

  game := models.Game {
    Name: "Pacman",
    //Id: vars["id"],
  }

  //GetAll("books");

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(game); err != nil {
      panic(err)
  }
}
