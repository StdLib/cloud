package route

import "net/http"

type router struct{}

var Router *router

func (*router) Mux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})
	return mux
}
