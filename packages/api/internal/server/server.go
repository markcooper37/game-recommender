package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/markcooper37/game-recommender/packages/api/internal/config"
	"github.com/markcooper37/game-recommender/packages/api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(conf config.Config) {
	db, err := NewDatabase(conf)
	if err != nil {
		log.Fatal("Could not start database")
	}

	if err = db.AutoMigrate(
		&models.Game{},
	); err != nil {
		log.Fatal("Could not migrate models")
	}

	handlers := mustInitRoutes(db, conf)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Handler: handlers,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func NewDatabase(conf config.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.New(postgres.Config{
		DSN: conf.DSN,
	}), &gorm.Config{})
}

func mustInitRoutes(db *gorm.DB, conf config.Config) http.Handler {
	schema, err := MakeGraphQLSchema(db)
	if err != nil {
		log.Fatalf("error creating routes %v", err)
	}

	handlers, err := Routes(schema, conf)
	if err != nil {
		log.Fatalf("error creating routes %v", err)
	}

	return handlers
}
