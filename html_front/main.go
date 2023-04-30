package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {
		var c Content = get_content()
		return_content(w, c)
	}

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Content struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Date    string `json:"date"`
	Content string `json:"content"`
}

func return_content(w http.ResponseWriter, content Content) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	t, err := template.ParseFiles("templates/default.html")
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(w, content)
	if err != nil {
		log.Fatal(err)
	}

}

func get_content() Content {
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://backend:9090", nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	response_json, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var content Content
	json.Unmarshal(response_json, &content)
	if err != nil {
		log.Fatal(err)
	}

	return content
}
