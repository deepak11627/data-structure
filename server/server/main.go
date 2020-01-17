package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print(fmt.Sprintf("Recieved request on %s %s", r.URL.Path, time.Now()))
	fmt.Fprint(w, fmt.Sprintf("You requested page %s", r.URL.Path))
}
