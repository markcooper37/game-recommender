package resolvers

import (
	"sort"

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
	games, err := g.Games()
	if err != nil {
		return nil, err
	}
	scores := make([]int, len(games))
	for index, game := range games {
		if args.Input.MinAge >= game.MinAge() {
			scores[index]++
		}
		if args.Input.Players >= game.MinPlayer() && args.Input.Players <= game.MaxPlayer() {
			scores[index]++
		}
		if args.Input.Category == game.Category() {
			scores[index]++
		}
		if args.Input.Genre == game.Genre() {
			scores[index]++
		}
	}

	sort.Slice(games, func(i, j int) bool {
		return scores[i] > scores[j]
	})

	return games, nil
}
