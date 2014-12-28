package main

import (
	"log"
	"net/http"
)

const VERSION = "0.1.0"

func main() {
	log.Printf("resolutionizerd %s starting...", VERSION)

	log.Fatalln(http.ListenAndServe(":8080", http.FileServer(http.Dir("./client"))))
}
