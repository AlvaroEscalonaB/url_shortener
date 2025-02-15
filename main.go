package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"

	"url_shortener/views"
)

func main() {
	component := views.Index()

	//mux := http.NewServeMux()

	// mux.HandleFunc("GET /static/", view.ServeStaticFiles)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
