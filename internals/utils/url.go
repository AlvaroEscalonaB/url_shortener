package utils

import (
	"fmt"
	"net/http"
)

func UrlReferer(w http.ResponseWriter, r *http.Request) string {
	scheme := "http"
	// Check if HTTPS is used
	if r.TLS != nil {
		scheme = "https"
	}

	fullURL := fmt.Sprintf("%s://%s", scheme, r.Host)
	return fullURL
}
