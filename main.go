package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

const VERSION = "0.1.0"

var clientDir string

func init() {
	clientEnv := os.Getenv("CLIENT")
	flag.StringVar(&clientDir, "client", clientEnv, "the directory where the client data is stored")
}

func main() {
	flag.Parse()
	fmt.Printf("resolutionizerd %s starting...", VERSION)
	fmt.Printf("listening on port %s", os.Getenv("PORT"))
	fmt.Printf("client root: %s", clientDir)

	fmt.Fatalln(http.ListenAndServe(":"+os.Getenv("PORT"), http.FileServer(http.Dir(clientDir))))
}
