package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

const VERSION = "0.1.0"

var clientDir string

func init() {
	clientEnv = os.Getenv("CLIENT")
	flag.StringVar(&clientDir, "client", clientDir, "the directory where the client data is stored")
}

func main() {
	flag.Parse()
	log.Printf("resolutionizerd %s starting...", VERSION)

	log.Fatalln(http.ListenAndServe(":"+os.Getenv("PORT"), http.FileServer(http.Dir(clientDir))))
}
