package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/markcooper37/game-recommender/packages/api/internal/config"
)

func StartServer(conf *config.Config) {
	http.HandleFunc("/", handleHome())
	http.HandleFunc("/allgames", handleAllGames())

	if err := http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil); err != nil {
		log.Fatal(err)
	}
}
