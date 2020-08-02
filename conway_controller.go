package main

import (
	"github.com/cadsdanaa/conwayHttp/universe"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/conway", conwayClosure())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func conwayClosure() func(w http.ResponseWriter, r *http.Request) {
	initialUniverse := universe.InitialUniverse(20, 1235412)
	return func(w http.ResponseWriter, r *http.Request) {
		initialUniverse.Progress()
		_, err := w.Write([]byte(universe.Draw(initialUniverse)))
		if err != nil {
			w.WriteHeader(500)
			log.Fatal(err)
		}
	}
}
