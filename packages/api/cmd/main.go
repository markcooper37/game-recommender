package main

import (
	"log"

	"github.com/markcooper37/game-recommender/packages/api/internal/config"
	"github.com/markcooper37/game-recommender/packages/api/internal/server"
)

func main() {
	conf := mustLoadConfig()
	server.StartServer(conf)
}

func mustLoadConfig() *config.Config {
	conf, err := config.Load()
	if err != nil {
		log.Fatalf("error loading app config %v", err)
	}
	return conf
}
