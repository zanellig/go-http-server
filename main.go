package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HttpHandler)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}

func HttpHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from: %s\n", r.URL)
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