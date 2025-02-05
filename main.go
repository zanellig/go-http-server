package main

import (
	"fmt"
	"http-server/db"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.Init()
	r := mux.NewRouter()
	r.HandleFunc("/transcription/{uuid}", TranscriptionHandler).Methods("GET").Schemes("http")
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Use the router in the second parameter
	http.ListenAndServe(":8080", r)
}

func TranscriptionHandler (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "Requested transcription with UUID: "+vars["uuid"])
}

/*
	If we want to control TLS, keep-alives and compression, should create a Transport
	*/
// func test(){

// 	tr := &http.Transport{
// 		MaxIdleConns:       10,
// 		IdleConnTimeout:    30 * time.Second,
// 		DisableCompression: true,
// 	}
// 	client := &http.Client{Transport: tr}
// 	resp, err := client.Get("https://example.com")
	
// }

/*
	Clients and Transports are safe for concurrent use by multiple goroutines and for efficiency should only be created once and re-used.
*/