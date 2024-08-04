package main

import (
	"log"
	"net/http"
)

func run() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})
	return http.ListenAndServe(":10443", nil)
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
