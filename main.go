package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"

	"url_shortener/internals/controllers"
	"url_shortener/internals/database"
	internal_view "url_shortener/internals/view"
	"url_shortener/views"
)

func main() {
	err := database.Db.CreateDatabase()
	if err != nil {
		fmt.Println("Error creating the database", err)
	}

	mux := http.NewServeMux()

	// mux.HandleFunc("GET /static/", view.ServeStaticFiles)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	mux.Handle("/", templ.Handler(views.Index()))
	mux.HandleFunc("POST /short-url", controllers.PostUrlToShorten)
	mux.HandleFunc("GET /favicon.ico", internal_view.FaviconContent)
	mux.HandleFunc("GET /{id}", controllers.RedirectShortUrl)

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", mux)
}
