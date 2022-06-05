package resolvers

import (
	"github.com/markcooper37/game-recommender/packages/api/internal/models"
	"gorm.io/gorm"
)

type Resolver struct {
	DB *gorm.DB
}

func (r *Resolver) Migrate() error {
	return r.DB.AutoMigrate(
		&models.Game{},
	)
}
