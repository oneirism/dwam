package main

import (
	"net/http"

	"github.com/oneirism/dwam/v2/internal/routes"
)

func main() {
	r := routes.NewDefaultRouter()

	r.WithGet("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	http.ListenAndServe("127.0.0.1:3333", r.GetMux())
}
