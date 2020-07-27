package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/conway", conwayHandler)
	http.ListenAndServe("localhost:8080", nil)
}

func conwayHandler(w http.ResponseWriter, r *http.Request) {
	_, e := w.Write([]byte("Conway Http"))
	if e != nil {
		w.WriteHeader(500)
	}
}
