package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Heya there!</h1>"))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)

}
