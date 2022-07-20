package resolvers

import (
	"github.com/markcooper37/game-recommender/packages/api/internal/models"
	"gorm.io/gorm"
)

type Game struct {
	game *models.Game
}

func (g *Game) Name() string {
	return g.game.Name
}

func (g *Game) MinAge() int32 {
	return g.game.MinAge
}

func (g *Game) MinPlayer() int32 {
	return g.game.MinPlayer
}

func (g *Game) MaxPlayer() int32 {
	return g.game.MaxPlayer
}

func (g *Game) Category() string {
	return g.game.Category
}

func (g *Game) Genre() string {
	return g.game.Genre
}

func (g *Game) Description() string {
	return g.game.Description
}

type GameResolver struct {
	GameQuery
	GameMutation
}

func NewGameResolver(db *gorm.DB) *GameResolver {
	return &GameResolver{
		NewGameQuery(db),
		NewGameMutation(db),
	}
}
