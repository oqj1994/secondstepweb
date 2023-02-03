package handlers

import (
	"fmt"
	"github.com/justinas/nosurf"
	"github.com/vitaLemoTea/secondstepweb/pkg/config"
	"github.com/vitaLemoTea/secondstepweb/pkg/model"
	"github.com/vitaLemoTea/secondstepweb/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	CF *config.Config
}

func NewHandler(r *Repository) {
	Repo = r
}

func NewRepo(cf *config.Config) *Repository {
	return &Repository{CF: cf}
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	msgs := map[string]string{
		"Greet": "hello，this is my question : old thing can not support our life",
	}
	d := model.TemplateData{StringMap: msgs}
	//*********
	{
		repo.CF.Session.Put(r.Context(), "cat", "kitty")
		repo.CF.Session.Put(r.Context(), "dog", "alex")

	}
	//*********
	d.CSRFToken = nosurf.Token(r)
	err := render.Render(w, "home.page.html", &d)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	msgs := map[string]string{
		"Greet": "hello，this is my question : old thing can not support our life",
	}
	d := model.TemplateData{StringMap: msgs}
	//*********
	{
		get, ok := repo.CF.Session.Get(r.Context(), "cat").(string)
		if !ok {
			d.StringMap["cat"] = "nothing"
		} else {
			d.StringMap["cat"] = get
		}
		get, ok = repo.CF.Session.Get(r.Context(), "dog").(string)
		if !ok {
			d.StringMap["dog"] = "nothing"
		} else {
			d.StringMap["dog"] = get
		}

	}
	//*********
	d.CSRFToken = nosurf.Token(r)
	err := render.Render(w, "about.page.html", &d)
	if err != nil {
		fmt.Println(err)
		return
	}

}
