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

	output := make([]*Game, 0, len(games))
	for _, game := range games {
		output = append(output, &Game{game: game})
	}

	return output, nil
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
	var games []*models.Game

	err := g.db.Find(&games).Error
	if err != nil {
		return nil, err
	}

	output := make([]*Game, 0, len(games))
	for _, game := range games {
		output = append(output, &Game{game: game})
	}
	
	scores := make([]int, len(output))
	for index, game := range output {
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

	sort.Slice(output, func(i, j int) bool {
		return scores[i] > scores[j]
	})

	return output, nil
}
