package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type PageData struct {
	Title string
}

var templates = template.Must(
	template.ParseFiles(
		"public/index.html",
	),
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("[ERROR] Unable to load environment variables: %s\n", err.Error())
	}
}

func main() {
	port := os.Getenv("PORT")
	address := os.Getenv("ADDRESS")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pageData := PageData{
			Title: "Home",
		}
		err := templates.ExecuteTemplate(w, "index.html", pageData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fs := http.FileServer(http.Dir("./public/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Printf("Starting server on http://%s:%s\n", address, port)
	err := http.ListenAndServe(address+":"+port, nil)
	if err != nil {
		log.Fatalf("[ERROR] unable to start server: %s\n", err.Error())
	}
}
