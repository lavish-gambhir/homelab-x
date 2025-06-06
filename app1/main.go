package main

import (
	"fmt"
	"log"
	"net/http"
)

const _port = "8080"

func healthz(w http.ResponseWriter, _ *http.Request) {
	log.Println("healthz::invoked")
	w.WriteHeader(http.StatusOK)
}

func readyz(w http.ResponseWriter, _ *http.Request) {
	// Test the core dependencies on the service functionality depends on
	// like db conn. etc.
	log.Println("readyz::invoked")
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello homelab")
	})
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/readyz", readyz)

	fmt.Println("taco listening on:", _port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", _port), nil); err != nil {
		log.Fatalf("err while running server: %v", err)
	}
}
