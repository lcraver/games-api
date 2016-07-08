package handlers

import (
  "encoding/json"
  "net/http"

  "github.com/lcraver/games-api/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
  index := models.Index {
    Version: "1.0.0",
  }

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(index); err != nil {
      panic(err)
  }
}
