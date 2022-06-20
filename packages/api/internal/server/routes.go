package server

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"strings"

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
		http.Error(w, fmt.Sprintf("playground handler error: %v", err), http.StatusInternalServerError)
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
	rootResolver := resolvers.NewRootResolver(db)

	return graphql.ParseSchema(schemaDefinition, rootResolver, opts...)
}

func ConcatGraphQLSchemaFiles(fsys fs.FS) (string, error) {
	var buf bytes.Buffer

	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("walking directory: %w", err)
		}

		fileSuffix := fmt.Sprintf(".%s", "graphql")
		if !strings.HasSuffix(path, fileSuffix) {
			return nil
		}

		f, err := fsys.Open(path)
		if err != nil {
			return fmt.Errorf("opening file %q: %w", path, err)
		}
		defer f.Close()

		fileInfo, err := f.Stat()
		if err != nil {
			return fmt.Errorf("getting stats for file %q: %w", path, err)
		}

		if fileInfo.IsDir() {
			return nil
		}

		_, err = io.Copy(&buf, f)
		if err != nil {
			return fmt.Errorf("writing %q bytes to buffer: %w", path, err)
		}

		_, err = fmt.Fprint(&buf, "\n")
		if err != nil {
			return fmt.Errorf("writing newline to buffer: %w", err)
		}

		return nil
	})
	if err != nil {
		return buf.String(), fmt.Errorf("walking content directory: %w", err)
	}

	return buf.String(), nil
}
