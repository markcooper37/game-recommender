package server

import (
	"encoding/json"
	"net/http"

	"github.com/markcooper37/game-recommender/packages/api/internal/models"
)

var Games = []models.Game{{Name: "Uno"}, {Name: "Monopoly"}, {Name: "Twister"}}

func handleHome() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Game Recommender!"))
	})
}

func handleAllGames() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Games)
	})
}
