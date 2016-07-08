package models

type Game struct {
  Name   string   `json:"name"`
  Id     string   `json:"id"`
}

type Games []Game
