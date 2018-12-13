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
	user     string
	pass     string
	realm    string
)

func init() {
	flag.StringVar(&localDir, "path", ".", "Path to serve from")
	flag.StringVar(&host, "host", "0.0.0.0", "Host to listen on")
	flag.IntVar(&port, "port", 8000, "Port to listen on")
	flag.StringVar(&user, "username", "", "Username for basic auth")
	flag.StringVar(&pass, "password", "", "Password for basic auth")
	flag.StringVar(&realm, "realm", "DUMB-HTTP", "Realm for basic auth")
}

func main() {
	flag.Parse()
	// Start serving...
	serve(fmt.Sprintf("%s:%d", host, port))
}

func serve(addr string) {
	mux := http.NewServeMux()
	// mux.HandleFunc("/testauth", SimpleAuth)
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
