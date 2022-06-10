package resolvers

type RootResolver struct {
	*GameResolver
}

func NewRootResolver() *RootResolver {
	return &RootResolver{
		GameResolver: newGameResolver(),
	}
}