package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var err error
	_, err = NewNode("https://mainnet.infura.io")
	if err != nil {
		log.Fatalf("Failed to init node: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "plop")
	})

	log.Printf("Listening on port %s\n\n", port)
	http.ListenAndServe(":"+port, nil)
}
