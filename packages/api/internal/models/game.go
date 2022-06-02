package models

type Game struct {
	Name        string `json:"Name"`
	MinAge      int    `json:"minAge"`
	MinPlayer   int    `json:"minPlayer"`
	MaxPlayer   int    `json:"maxPlayer"`
	Category    string `json:"category"`
	Description string `json:"description"`
}
