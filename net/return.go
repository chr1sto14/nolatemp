package net

import (
	"encoding/json"
	"net/http"
)

func Bad(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}

func Good(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func Json(w http.ResponseWriter, jsn interface{}) {
	Good(w)
	w.Header().Set("Content-Type", "application/json")
	buf, _ := json.Marshal(jsn)
	w.Write(buf)
}
