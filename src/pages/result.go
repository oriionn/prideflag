package pages

import (
	"context"
	_ "embed"
	"html/template"
	"net/http"

	"gorm.io/gorm"
	"prideflag.fun/src/data"
	"prideflag.fun/src/database"
)

//go:embed templates/result.html
var resultContent string

type ResultPageData struct {
	Test 	database.Test
	Message string
	Max int
}

func Result(db *gorm.DB, ctx context.Context) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		if !query.Has("t") {
			http.Error(w, "These results aren't available", 404)
			return
		}

		id := query.Get("t")
		test, err := gorm.G[database.Test](db).Where("id = ?", id).First(ctx)
		if err != nil {
			http.Error(w, "These results aren't available", 404)
			return
		}

		message := "You can do better!"
		if test.Note >= len(data.DATASET) - 2 {
			message = "You are a fucking true faggot :3"
		} else if test.Note >= len(data.DATASET) / 2 {
			message = "Good job!"
		} else if test.Note <= 2 {
			message = "HOMOPHOBIC!!!! :("
		}

		data := ResultPageData{
			Test: test,
			Message: message,
			Max: len(data.DATASET),
		}

		// Templating
		t, err := template.New("results").Parse(resultContent)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		t.Execute(w, data)
	}
}
