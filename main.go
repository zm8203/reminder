package main

import (
	"awesomeProject/router"
	"log"
	"net/http"
)

func main() {
	r := router.SetupRouter()

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
