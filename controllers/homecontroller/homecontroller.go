package homecontroller

import (
	"KnapSack/models/makananmodel"
	"KnapSack/views"
	"html/template"
	"net/http"
	"strconv"
)

var tmpl = template.Must(template.ParseFS(views.FS, "home/*.html"))

func Splash(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func ShowGenerated(w http.ResponseWriter, r *http.Request) {
	frequency, _ := strconv.Atoi(r.FormValue("frequency"))
	budget, _ := strconv.ParseFloat(r.FormValue("budget"), 64)

	dt_makanan := makananmodel.GenerateData(frequency, budget)
	data := map[string]any{
		"dt_makanan": dt_makanan,
	}

	tmpl.ExecuteTemplate(w, "generatedData.html", data)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "home.html", nil)
}