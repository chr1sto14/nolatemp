package net

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func SendJson(url string, pkg interface{}) (err error) {
	buf, err := json.Marshal(pkg)
	if err != nil {
		return
	}

	r, err := http.Post(url, "application/json", bytes.NewReader(buf))
	if err != nil {
		return
	}
	defer r.Body.Close()
	return
}
