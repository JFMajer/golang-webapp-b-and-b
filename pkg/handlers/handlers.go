package handlers

import (
	"net/http"

	"github.com/JFMajer/golang-webapp-b-and-b/pkg/config"
	"github.com/JFMajer/golang-webapp-b-and-b/pkg/models"
	"github.com/JFMajer/golang-webapp-b-and-b/pkg/render"
)

// Repo the repositoruy used by the handlers
var Repo *Repository

// Repository is a repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, passing data to template"
	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
