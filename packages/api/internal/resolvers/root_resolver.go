package resolvers

import "gorm.io/gorm"

type RootResolver struct {
	*GameResolver
}

func NewRootResolver(db *gorm.DB) *RootResolver {
	return &RootResolver{
		GameResolver: NewGameResolver(db),
	}
}
