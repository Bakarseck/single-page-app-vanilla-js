package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	LoadEnv(".env")

	fs := http.FileServer(http.Dir("./frontend/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend/index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Démarrage du serveur
	log.Printf("Server running on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
