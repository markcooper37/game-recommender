package models

type Game struct {
	Name        string `gorm:"primaryKey"`
	MinAge      int32
	MinPlayer   int32
	MaxPlayer   int32
	Category    string // card, board, video
	Genre       string // adult, trivia, action
	Description string
}
