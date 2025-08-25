package main

import (
	"log"
	"net/http"
)

// First we define a handler function for the default landing url (/)
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	mux := http.NewServeMux() // initialize a new servemux

	mux.HandleFunc("/", home) // register the home function as the handler for the "/" page

	log.Print("Starting server on :4000")

	// Use this function to start a web server. We pass two parameters, TCP network address to listen on
	// and the servemux we created before.
	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}
