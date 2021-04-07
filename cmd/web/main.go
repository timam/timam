package main

import (
	"fmt"
	"github.com/timam/timam/pkg/config"
	"github.com/timam/timam/pkg/handlers"
	"github.com/timam/timam/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"


// main is th main application function
func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil{
		log.Fatal("cant render template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
