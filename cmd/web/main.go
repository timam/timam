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

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
}
