package models

type Game struct {
	Name        string
	MinAge      int32
	MinPlayer   int32
	MaxPlayer   int32
	Category    string
	Description string
}
