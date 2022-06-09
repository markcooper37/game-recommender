package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/markcooper37/game-recommender/packages/api/internal/config"
	"github.com/markcooper37/game-recommender/packages/api/internal/resolvers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(conf config.Config) {
	db, err := NewDatabase(conf)
	if err != nil {
		log.Fatal("Could not start database")
	}

	resolver := resolvers.Resolver{DB: db}
	if err = resolver.Migrate(); err != nil {
		log.Fatal("Could not migrate models")
	}

	http.HandleFunc("/", handleHome())

	if err := http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), nil); err != nil {
		log.Fatal(err)
	}
}

func NewDatabase(conf config.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.New(postgres.Config{
		DSN: conf.DSN,
	}), &gorm.Config{})
}
