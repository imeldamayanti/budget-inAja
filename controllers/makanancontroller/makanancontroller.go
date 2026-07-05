package makanancontroller

import (
	"KnapSack/models/makananmodel"
	"KnapSack/views"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	dt_makanan := makananmodel.GetAll()
	data := map[string]any{
		"dt_makanan": dt_makanan,
	}

	temp, err := template.ParseFS(views.FS, "home/data.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}