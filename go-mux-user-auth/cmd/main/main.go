package main

import (
	"log"
	"net/http"

	"github.com/KannanPalani57/go-mux-user-auth/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterUserRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8081", r))
}
