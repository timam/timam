package main

import (
	"fmt"
	"github.com/timam/timam/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"


// main is th main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
