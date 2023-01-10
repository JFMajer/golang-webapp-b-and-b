package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JFMajer/golang-webapp-b-and-b/pkg/config"
	"github.com/JFMajer/golang-webapp-b-and-b/pkg/handlers"
	"github.com/JFMajer/golang-webapp-b-and-b/pkg/render"
)

const portNumber = ":8080"

// main is main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	// http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
