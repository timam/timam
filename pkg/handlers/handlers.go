package handlers

import (
	"github.com/timam/timam/pkg/config"
	"github.com/timam/timam/pkg/render"
	"net/http"
)

//Repo the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository  {
	return &Repository{
		App : a,
	}
}

//NewHandlers sets the repository for handlers
func NewHandlers(r *Repository)  {
	Repo = r
}

//Home is the home page handlers
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gohtml")
}

//About is the about page handlers
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.gohtml")
}

