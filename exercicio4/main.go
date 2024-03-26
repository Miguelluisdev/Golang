package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		go func() {
			fmt.Fprintf(w, "hello world um novo servidor criado")

		}()
	})
	http.ListenAndServe(":8080", mux)
}
