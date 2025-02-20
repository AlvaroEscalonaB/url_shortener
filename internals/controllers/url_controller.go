package controllers

import (
	"fmt"
	"net/http"
	"url_shortener/internals/database"
	"url_shortener/internals/utils"
	"url_shortener/views/components"
	"url_shortener/views/layouts"
)

func PostUrlToShorten(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	url := r.Form.Get("url")

	if url == "" {
		component := components.TransformedUrl("Cannot generate url")
		component.Render(r.Context(), w)
		return
	}
	fmt.Println("url", url)
	newShortUrl, err := database.Db.CreateShortUrl(url)

	if err != nil {
		component := components.TransformedUrl("Cannot parse and generate new url, try again.")
		component.Render(r.Context(), w)
		return
	}

	referer := utils.UrlReferer(w, r)
	refererShortUrl := fmt.Sprintf("%s/%s", referer, newShortUrl.ShortUrl)

	component := components.TransformedUrl(refererShortUrl)
	component.Render(r.Context(), w)
}

func RedirectShortUrl(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("id")
	if pathValue == "" {
		component := layouts.Layout()
		component.Render(r.Context(), w)
		return
	}

	urlRecord, err := database.Db.QueryShortUrl(pathValue)
	fmt.Println("record url", urlRecord.Url)
	if err != nil {
		fmt.Println("ERROR")
		component := layouts.NotFound()
		component.Render(r.Context(), w)
		return
	}
	http.Redirect(w, r, urlRecord.Url, http.StatusFound)
}
