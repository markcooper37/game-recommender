package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/markcooper37/game-recommender/packages/api/internal/models"
)

var Games = []models.Game{{Name: "Uno"}, {Name: "Monopoly"}, {Name: "Twister"}}

func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Game Recommender!"))
	})

	http.HandleFunc("/allgames", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Games)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
