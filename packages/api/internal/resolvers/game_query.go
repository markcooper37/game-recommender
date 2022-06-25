package resolvers

import (
	"github.com/markcooper37/game-recommender/packages/api/internal/models"
	"gorm.io/gorm"
)

type GameQuery struct {
	db *gorm.DB
}

func NewGameQuery(db *gorm.DB) GameQuery {
	return GameQuery{db: db}
}

func (g *GameQuery) Games() ([]*Game, error) {
	var games []*models.Game
	err := g.db.Find(&games).Error
	if err != nil {
		return nil, err
	}

	resp := make([]*Game, 0, len(games))
	for _, game := range games {
		resp = append(resp, &Game{game: game})
	}
	return resp, nil
}

type searchGamesArgs struct {
	Input SearchGamesInput
}

type SearchGamesInput struct {
	MinAge int32
	Players int32
	Category string
	Genre string
}

func (g *GameQuery) SearchGames(args searchGamesArgs) ([]*Game, error) {
	// Give each game points based on how well matched they are to the input criteria
	// Order the games based on their points with the best matches first
	return []*Game{}, nil
}
