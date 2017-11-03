package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "strings"
)

func main() {
	http.HandleFunc("/", testHandler)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}

type UrlMeta struct {
	Status        int    `json:"status"`
	ContentType   string `json:"content-type"`
	ContentLength int    `json:"content-length"`
}

type UrlDataElement struct {
	TagName string `json:"tag-name"`
	Count   int    `json:"count"`
}

type UrlData struct {
	Url      string           `json:"url"`
	Meta     UrlMeta          `json:"meta"`
	Elements []UrlDataElement `json:"elements"`
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	var data []UrlData
	var input []string

	err := decoder.Decode(&input)
	if err != nil {
		encoder.Encode(err)
	}

	for _, url := range input {
		data = append(data, getData(url))
	}

	encoder.Encode(data)
}

func getData(url string) UrlData {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	_ = resp

	meta := UrlMeta{
		Status:        resp.StatusCode,
		ContentType:   resp.Header.Get("Content-type"),
		ContentLength: 1,
	}

	data := UrlData{Url: url, Meta: meta}

	return data
}
