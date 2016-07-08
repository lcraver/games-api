package models

type Game struct {
  Id          int     `json:"id"`
  Title       string  `json:"title"`
  Developer   string   `json:developer`
  Publisher   string   `json:publisher`
  Languages   string   `json:languages`
}

type Games []Game
