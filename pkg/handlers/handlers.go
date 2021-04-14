package handlers

import (
	"github.com/timam/timam/pkg/config"
	"github.com/timam/timam/pkg/modles"
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
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.gohtml", &modles.TemplateData{})
}

//About is the about page handlers
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.gohtml", &modles.TemplateData{
		StringMap: stringMap,
	})
}

