package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"

	"url_shortener/internals/database"
	"url_shortener/internals/utils"
	"url_shortener/views"
	"url_shortener/views/components"
)

func PostUrlToShorten(w http.ResponseWriter, r *http.Request) {
	url := r.Form.Get("url")

	if url != "" {
		component := components.TransformedUrl("Cannot generate url")
		component.Render(r.Context(), w)
		return
	}

	fmt.Println("Form data related to url", url)

	newShortUrl, err := database.Db.CreateShortUrl(url)
	if err != nil {
		component := components.TransformedUrl("Cannot parse and generate new url, try again.")
		component.Render(r.Context(), w)
		return
	}

	referer := utils.UrlReferer(w, r)
	refererShortUrl := fmt.Sprintf("%s/short-url/%s", referer, newShortUrl.ShortUrl)

	component := components.TransformedUrl(refererShortUrl)
	component.Render(r.Context(), w)
}

func main() {
	err := database.Db.CreateDatabase()
	if err != nil {
		fmt.Println("Error creating the database", err)
	}

	mux := http.NewServeMux()

	// mux.HandleFunc("GET /static/", view.ServeStaticFiles)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	mux.Handle("/", templ.Handler(views.Index()))
	mux.HandleFunc("POST /short-url", PostUrlToShorten)
	// http.Handle("/", )

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", mux)
}
