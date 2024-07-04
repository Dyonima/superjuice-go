package views

import (
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method + " " + r.URL.Path)
	w.Write([]byte("Hello World!"))
}
