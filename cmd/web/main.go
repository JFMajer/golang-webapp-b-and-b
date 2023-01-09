package main

import (
	"fmt"
	"net/http"

	"github.com/JFMajer/golang-webapp-b-and-b/pkg/handlers"
)

const portNumber = ":8080"

// main is main application function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)

}
