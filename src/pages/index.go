package pages

import (
	_ "embed"
	"html/template"
	"net/http"
)

//go:embed templates/index.html
var indexContent string

type IndexPageData struct {
	IsFinished bool
}

func Index(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("finished")
	isFinished := err != nil

	if err == nil {
		isFinished = cookie.Value == "true"
	}

	// Templating
	t, err := template.New("index").Parse(indexContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := IndexPageData{
		IsFinished: isFinished,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(w, data)
}
