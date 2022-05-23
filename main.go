package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "ok")
		fmt.Printf("ping, %v\n", time.Now().String())
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
