package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello homelab")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("err while running server: %v", err)
	}
}
