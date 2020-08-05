package main

import (
	"github.com/cadsdanaa/conwayHttp/universe"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/conway", conwayClosure())
	port := "5000"
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
		if !authenticated(r, w) {
			return
		}
		initialUniverse.Progress()
		_, err := w.Write([]byte(universe.Draw(initialUniverse)))
		if err != nil {
			w.WriteHeader(500)
			log.Fatal(err)
		}
	}
}

func authenticated(r *http.Request, w http.ResponseWriter) bool {
	user, pass, _ := r.BasicAuth()
	if os.Getenv("CONWAY_USER") != user || os.Getenv("CONWAY_PASSWORD") != pass {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return false
	}
	return true
}
