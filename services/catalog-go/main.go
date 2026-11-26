package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	server := &http.Server{
		Addr:              ":" + port,
		Handler:           newRouter(),
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("catalog-go listening on %s", port)
	log.Fatal(server.ListenAndServe())
}
