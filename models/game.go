package models

type Game struct {
  Id     int      `json:"id"`
  Name   string   `json:"name"`
  Price  int      `json:price`
}

type Games []Game
