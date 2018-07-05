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
	port     = "8000"
)

func init() {
	flag.StringVar(&localDir, "path", ".", "Path to serve from")
}

func main() {
	flag.Parse()
	// Check for port in commandline
	if commandPort := flag.Arg(0); commandPort != "" {
		port = commandPort
	}
	// Start serving...
	serve(host + ":" + port)
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
