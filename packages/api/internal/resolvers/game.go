package resolvers

type GameResolver struct{}

func newGameResolver() *GameResolver {
	return &GameResolver{}
}

func (h *GameResolver) GameRecommenderString() string {
	return "Welcome to Game Recommender!"
}