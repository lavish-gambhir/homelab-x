package main

import (
	"fmt"
	"log"
	"net/http"
)

const _port = "8080"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello homelab")
	})

	fmt.Println("taco listening on:", _port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", _port), nil); err != nil {
		log.Fatalf("err while running server: %v", err)
	}
}
