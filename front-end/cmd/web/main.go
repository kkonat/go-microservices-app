package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Render(w, "test.page.gohtml")
	})
	// http.HandleFunc("/health-check", HealthCheckHandler)

	fmt.Println("Starting front end service on port 80")
	err := http.ListenAndServe("127.0.0.1:80", nil)
	if err != nil {
		log.Panic(err)
	}
}
