package server

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

func ConcatGraphQLSchemaFiles(fsys fs.FS) (string, error) {
	var buf bytes.Buffer

	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("walking dir: %w", err)
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