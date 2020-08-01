package main

import (
	"github.com/cadsdanaa/conwayHttp/universe"
	"net/http"
)

func main() {
	http.HandleFunc("/conway", conwayClosure())
	http.ListenAndServe("localhost:8080", nil)
}

func conwayClosure() func(w http.ResponseWriter, r *http.Request) {
	initialUniverse := universe.InitialUniverse(30, 1235412)
	return func(w http.ResponseWriter, r *http.Request) {
		initialUniverse.Progress()
		_, e := w.Write([]byte(universe.Draw(initialUniverse)))
		if e != nil {
			w.WriteHeader(500)
		}
	}
}
