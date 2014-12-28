package main

import (
	"flag"
	"fmt"
	"io"
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

	if clientDir == "" {
		clientDir = os.Getenv("CLIENT")
	}

	fmt.Printf("client root: %s\n", clientDir)

	if _, err := os.Stat(clientDir); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	http.Handle("/", LoggingHandler(os.Stdout, http.FileServer(http.Dir(clientDir))))

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type loggingHandler struct {
	writer  io.Writer
	handler http.Handler
}

func (h loggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(h.writer, "%s %-50s %s\n", r.Method, r.RequestURI, r.Header.Get("User-Agent"))
	h.handler.ServeHTTP(w, r)
}

func LoggingHandler(w io.Writer, h http.Handler) http.Handler {
	return loggingHandler{w, h}
}
