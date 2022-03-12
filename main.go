package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TImestamp: %s", time.Now())
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Println("Server running at port 80")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func main() {
	handleRequests()
}
