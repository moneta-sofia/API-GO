package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, 200, "YEPIII!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
