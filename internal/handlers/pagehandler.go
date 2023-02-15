package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/justinas/nosurf"
	"github.com/vitaLemoTea/secondstepweb/internal/config"
	"github.com/vitaLemoTea/secondstepweb/internal/model"
	"github.com/vitaLemoTea/secondstepweb/internal/render"
	"log"
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
	////*********
	//{
	//	repo.CF.Session.Put(r.Context(), "cat", "kitty")
	//	repo.CF.Session.Put(r.Context(), "dog", "alex")
	//
	//}
	////*********
	err := render.Render(w, "home.page.html", r, &d)
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
	err := render.Render(w, "about.page.html", r, &d)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func (repo *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	data := model.TemplateData{}
	err := render.Render(w, "general.page.html", r, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (repo *Repository) Major(w http.ResponseWriter, r *http.Request) {
	data := model.TemplateData{}
	err := render.Render(w, "major.page.html", r, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (repo *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	data := model.TemplateData{}
	err := render.Render(w, "reservation.page.html", r, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (repo *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	data := model.TemplateData{}
	err := render.Render(w, "reservation.page.html", r, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (repo *Repository) Book(w http.ResponseWriter, r *http.Request) {
	data := model.TemplateData{}

	err := render.Render(w, "book.page.html", r, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (repo *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	s := fmt.Sprintf("start-time:%s :end-time:%s", start, end)
	w.Write([]byte(s))
	return
}

func (repo *Repository) AvailabilityJson(w http.ResponseWriter, r *http.Request) {
	t := struct {
		Message string `json:"message,omitempty"`
		OK      bool   `json:"ok,omitempty"`
	}{
		"you are so wonderful",
		true,
	}
	bs, err := json.MarshalIndent(t, "", " ")
	log.Println(string(bs))
	if err != nil {
		http.Error(w, "json error", http.StatusOK)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bs)
	return
}
