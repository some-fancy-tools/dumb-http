package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	localDir = "."
	host     = "0.0.0.0"
	port     = 8000
)

func init() {
	flag.StringVar(&localDir, "path", ".", "Path to serve from")
	flag.StringVar(&host, "host", "0.0.0.0", "Host to listen on")
	flag.IntVar(&port, "port", 8000, "Port to listen on")
}

func main() {
	flag.Parse()
	// Start serving...
	serve(fmt.Sprintf("%s:%d", host, port))
}

func serve(addr string) {
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(localDir))))
	server := http.Server{
		Addr:    addr,
		Handler: NewLoggingHandler(mux, os.Stdout),
	}
	fmt.Printf("Serving at http://%s/ from %s\n", addr, localDir)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
