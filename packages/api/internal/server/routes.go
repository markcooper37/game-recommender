package server

import (
	"embed"
	"net/http"

	upload "github.com/eko/graphql-go-upload"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/markcooper37/game-recommender/packages/api/internal/config"
	"github.com/markcooper37/game-recommender/packages/api/internal/resolvers"
	"go.opencensus.io/plugin/ochttp"
	"gorm.io/gorm"
)

//go:embed graphql/*.graphql
var folder embed.FS

func Routes(schema *graphql.Schema, conf config.Config) (http.Handler, error) {
	r := chi.NewRouter()

	r.Use(
		middleware.Recoverer,
		middleware.Heartbeat("/ping"),
	)

	r.Method(http.MethodPost, "/api/graphql", upload.Handler(&relay.Handler{Schema: schema}))

	return &ochttp.Handler{Handler: r}, nil
}

func MakeGraphQLSchema(db *gorm.DB) (*graphql.Schema, error) {
	schemaDefinition, err := ConcatGraphQLSchemaFiles(folder)
	if err != nil {
		return nil, err
	}
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	rootResolver := resolvers.NewRootResolver()

	return graphql.ParseSchema(schemaDefinition, rootResolver, opts...)
}
