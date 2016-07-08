package models

type Game struct {
  Name   string   `json:"name"`
  Id     int   `json:"id"`
}

type Games []Game
