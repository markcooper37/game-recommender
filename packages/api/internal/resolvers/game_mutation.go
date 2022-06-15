package resolvers

import (
	"github.com/markcooper37/game-recommender/packages/api/internal/models"
	"gorm.io/gorm"
)

type GameMutation struct {
	db *gorm.DB
}

func NewGameMutation(db *gorm.DB) GameMutation {
	return GameMutation{db: db}
}

type createGameArgs struct {
	Input *models.Game
}

func (g *GameMutation) CreateGame(args createGameArgs) (string, error) {
	return "ok", g.db.Create(args.Input).Error
}

type deleteGameArgs struct {
	Name string
}

func (g *GameMutation) DeleteGame(args deleteGameArgs) (string, error) {
	err := g.db.Delete(&models.Game{}, "name = ?", args.Name).Error
	if err != nil {
		return "", err
	}
	return "successfully deleted " + args.Name, nil
}
