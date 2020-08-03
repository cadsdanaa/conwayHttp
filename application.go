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
	size := 20
	var seed int64 = 1235412
	log.Printf("Creating initial universe of size %d with randomization seed %d", size, seed)
	initialUniverse := universe.InitialUniverse(size, seed)
	return func(w http.ResponseWriter, r *http.Request) {
		initialUniverse.Progress()
		_, err := w.Write([]byte(universe.Draw(initialUniverse)))
		if err != nil {
			w.WriteHeader(500)
			log.Fatal(err)
		}
	}
}
