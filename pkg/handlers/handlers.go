package handlers

import (
	"github.com/akashprabu/bookings/pkg/config"
	"github.com/akashprabu/bookings/pkg/models"
	"github.com/akashprabu/bookings/pkg/render"
	"net/http"
)

// Repository is the respository type
type Respository struct {
	app *config.AppConfig
}

// Repo the repository used by the handlers
var Repo *Respository

func NewRepo(a *config.AppConfig) *Respository {
	return &Respository{
		app: a,
	}
}

func NewHandlers(r *Respository) {
	Repo = r
}

func (repo *Respository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	repo.app.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplates(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the About Page of the application
func (repo *Respository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"
	remoteIp := repo.app.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp
	render.RenderTemplates(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
