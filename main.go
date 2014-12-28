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
	fmt.Printf("resolutionizerd %s starting...\n", VERSION)
	fmt.Printf("listening on port %s\n", os.Getenv("PORT"))
	fmt.Printf("client environment variable setting: %s\n", os.Getenv("CLIENT"))
	fmt.Printf("client root: %s\n", clientDir)

	if clientDir == "" {
		clientDir = os.Getenv("CLIENT")
	}

	if _, err := os.Stat(clientDir); err != nil {
		fmt.Println(err)
	}

	http.Handle("/", http.FileServer(http.Dir(clientDir)))

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
