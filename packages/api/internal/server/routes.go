package server

import (
	"embed"
	"fmt"
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

//go:embed templates/playground.html
var playgroundHTML []byte

func playground(w http.ResponseWriter, _ *http.Request) {
	if _, err := w.Write(playgroundHTML); err != nil {
		http.Error(w, fmt.Sprintf("playground handler error %v", err), http.StatusInternalServerError)
		return
	}
}

//go:embed graphql/*.graphql
var folder embed.FS

func Routes(schema *graphql.Schema, conf config.Config) http.Handler {
	r := chi.NewRouter()

	r.Use(
		middleware.Recoverer,
		middleware.Heartbeat("/ping"),
	)

	r.Get("/api/graphql", playground)

	r.Method(http.MethodPost, "/api/graphql", upload.Handler(&relay.Handler{Schema: schema}))

	return &ochttp.Handler{Handler: r}
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
