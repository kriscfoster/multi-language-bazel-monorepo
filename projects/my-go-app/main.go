package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "github.com/kriscfoster/multi-language-bazel-monorepo/projects/greeter"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("received request")
    w.Write([]byte(greeter.Greet()))
}

func main() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", YourHandler)

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))
}
