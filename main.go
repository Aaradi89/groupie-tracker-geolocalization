package main

import (
	"fmt"
	"groupie-tracker/backend"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", backend.HomePage)          // Handle Home Page
	http.HandleFunc("/Artist", backend.ArtistsPage) // Handle Artist page

	go func() { // To run server in background
		log.Fatal(http.ListenAndServe(":8080", nil)) // Run server
	}()
	fmt.Println("http://localhost:8080/")
	select {} // Keep program running
}
