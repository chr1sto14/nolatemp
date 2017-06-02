package net

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func SendJson(url string, pkg interface{}) error {
	// TODO json Encoder
	buf, _ := json.Marshal(pkg)
	r, err := http.Post(url, "application/json", bytes.NewReader(buf))
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return nil
}
