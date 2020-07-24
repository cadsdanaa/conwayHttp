package conwayHttp

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/conway", conwayHandler)
	http.ListenAndServe("localhost:8080", nil)
}

func conwayHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
