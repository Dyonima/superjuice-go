package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/KyleKincer/superjuice-go/internal/api"
	"github.com/KyleKincer/superjuice-go/internal/views"
	"github.com/a-h/templ"
)

const port = 8080

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method + " " + r.URL.Path)
	http.ServeFile(w, r, "./assets/favicon.ico")
}

func main() {
	http.Handle("/", templ.Handler(views.Home()))
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/api/calculate", api.CalculateHandler)

	log.Println("Server started on port " + strconv.Itoa(port) + "...")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
