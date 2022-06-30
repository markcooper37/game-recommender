package resolvers

import (
	"errors"

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
	if args.Input.MinPlayer > args.Input.MaxPlayer {
		return "", errors.New("minimum number of players must be smaller than maximum number of players")
	} else if args.Input.MinAge < 0 || args.Input.MinAge > 150 {
		return "", errors.New("invalid age")
	}
	
	err := g.db.Create(args.Input).Error
	if err != nil {
		return "", err
	}

	return "successfully added " + args.Input.Name, nil
}

type updateGameArgs struct {
	Input UpdateGameInput
}

type UpdateGameInput struct {
	Name        string
	MinAge      int32
	MinPlayer   int32
	MaxPlayer   int32
	Category    string
	Genre       string
	Description string
}

func (g *GameMutation) UpdateGame(args updateGameArgs) (*Game, error) {
	updates := models.Game{
		Name: args.Input.Name,
		MinAge: args.Input.MinAge,
		MinPlayer: args.Input.MinPlayer,
		MaxPlayer: args.Input.MaxPlayer,
		Category: args.Input.Category,
		Genre: args.Input.Genre,
		Description: args.Input.Description,
	}

	err := g.db.Model(&models.Game{}).Where("name = ?", args.Input.Name).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	return &Game{game: &updates}, nil
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
