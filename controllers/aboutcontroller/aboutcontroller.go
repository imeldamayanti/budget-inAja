package aboutcontroller

import (
	"html/template"
	"net/http"

	"KnapSack/views"
)

var tmpl = template.Must(template.ParseFS(views.FS, "home/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.html", nil)
}