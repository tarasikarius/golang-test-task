package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"regexp"
)

func main() {
	http.HandleFunc("/", mainHandler)
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

func mainHandler(w http.ResponseWriter, r *http.Request) {
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
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	elements := countTags(body)

	meta := UrlMeta{
		Status:        resp.StatusCode,
		ContentType:   resp.Header.Get("Content-type"),
		ContentLength: len([]rune(string(body))),
	}

	data := UrlData{Url: url, Meta: meta, Elements: elements}

	return data
}

func countTags(body []byte) []UrlDataElement {
	var elements []UrlDataElement
	tagsCount := make(map[string]int)

	reg := regexp.MustCompile("<([a-z]+)([^>]*)>")
	tags := reg.FindAll(body, -1);

	reg = regexp.MustCompile("([a-z]+)")
	for _, tag := range tags {
		tagName := string(reg.Find(tag))

		if tagsCount[tagName] == 0 {
			tagsCount[tagName] = 1
		} else {
			tagsCount[tagName]++
		}
	}

	for tag, count := range tagsCount {
		elements = append(elements, UrlDataElement{tag, count})
	}

	return elements
}