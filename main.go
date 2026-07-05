package main

import (
	"KnapSack/config"
	"KnapSack/assets"
	"KnapSack/controllers/aboutcontroller"
	"KnapSack/controllers/homecontroller"
	"KnapSack/controllers/makanancontroller"
	// "KnapSack/models/makananmodel"
	"log"
	"net/http"
	"os"
)

func main() {
	config.ConnectDB()

	fs := http.FileServer(http.FS(assets.FS))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/", homecontroller.Splash)
	http.HandleFunc("/generated-data", homecontroller.ShowGenerated)
	http.HandleFunc("/data", makanancontroller.Index)
	http.HandleFunc("/about", aboutcontroller.Index)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	log.Println("Server running on port " + port)
	http.ListenAndServe("0.0.0.0:"+port, nil)
}

