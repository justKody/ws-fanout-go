package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/ws", handleWs)

	http.ListenAndServe(":8080", router)

}

func handleWs(w http.ResponseWriter, r *http.Request) {
	hub := NewHub()
	go hub.Run()
	
	serveWs(hub, w, r)
}
