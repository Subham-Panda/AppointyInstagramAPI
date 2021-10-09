package utils

import (
	"errors"
	"net/http"
	"strings"
)

func GetIDFromURL(r *http.Request) (string, error) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		return "0", errors.New("no ID found")
	}
	id := parts[len(parts)-1]
	return id, nil
}

func GetUserIDFromPostURL(r *http.Request) (string, error) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 4 {
		return "0", errors.New("no ID found")
	}
	id := parts[len(parts)-1]
	return id, nil
}