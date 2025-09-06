package pages

import (
	"context"
	"embed"
	"fmt"
	"io"
	"mime"
	"net/http"
	"path/filepath"

	"gorm.io/gorm"
	"prideflag.fun/src/database"
)

func Flag(db *gorm.DB, ctx context.Context, public embed.FS) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		if !query.Has("f") {
			http.Error(w, "No query provided", http.StatusBadRequest)
			return
		}

		file := query.Get("f")
		image, err := gorm.G[database.Images](db).Where("id = ?", file).First(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		path := fmt.Sprintf("public/flags/%s", image.File)
		f, err := public.Open(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// Detect mime type
		ext := filepath.Ext(image.File)
		mimeType := mime.TypeByExtension(ext)
		if mimeType == "" {
			mimeType = "image/png"
		}

		w.Header().Set("Content-Type", mimeType)
		w.WriteHeader(http.StatusOK)
		io.Copy(w, f)
	}
}
