package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler a function which writes a byte slice containing
// "Hello from Snips" as the response body
func home(w http.ResponseWriter, r *http.Request) {
  // Check if the current request URL path exactly matches "/". If it doesn't,
  // use the http.NotFound() function to send a 404 response to the client.
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }
  w.Write([]byte("Hello from Snips"))
}

// Add a snipsView handler function
func snipsView(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    http.NotFound(w, r)
    return
  }

  fmt.Fprintf(w, "Display a specific snip with ID %d...", id)
}

// Add a snipCreate handler function
func snipCreate(w http.ResponseWriter, r *http.Request) {
  // Use r.Method to check whether the request is using POST or not.
  if r.Method != "POST" {
    w.Header().Set("Allow", http.MethodPost)
    // w.WriteHeader(405)
    // w.Write([]byte("Method Not Allowed"))
    http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    return
  }
  w.Write([]byte("Create a new snippet..."))
}

func main() {
  // Use the http.NewServeMux() function to initialize a new servermux, then
  // register the home function as the handler for the "/" URL pattern.
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snip/view", snipsView)
  mux.HandleFunc("/snip/create", snipCreate)

  log.Println("Starting server on :4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}
