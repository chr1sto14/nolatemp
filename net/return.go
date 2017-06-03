package net

import (
	"net/http"
)

func Bad(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}
