package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kriscfoster/multi-language-bazel-monorepo/projects/go_hello_world"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request")
	w.Write([]byte(go_hello_world.HelloWorld() + " Bazoku is cool!!"))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}

	return port
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)
	// Bind to a port and pass our router in
	port := getPort()
	log.Println("Going to listen on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
