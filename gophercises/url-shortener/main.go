package main

import (
	"fmt"
	"net/http"

	urlshort "github.com/gabriel-leone/studying-go/gophercises/url-shortener/urlshort"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	yamlData := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yamlHandler, err := urlshort.YAMLHandler([]byte(yamlData), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on: 8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
