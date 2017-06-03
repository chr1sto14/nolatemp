package net

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 20 * time.Second}

func ParseJson(body io.Reader, target interface{}) error {
	return json.NewDecoder(body).Decode(target)
}

func GetJson(url string, target interface{}) error {
	resp, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return ParseJson(resp.Body, target)
}
