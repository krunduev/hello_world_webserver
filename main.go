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
		fmt.Printf("%v, %v\n", r.RequestURI, time.Now().Format("2006-01-02 15:04:05"))
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
