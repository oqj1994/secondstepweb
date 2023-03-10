package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vitaLemoTea/secondstepweb/internal/config"
	"github.com/vitaLemoTea/secondstepweb/internal/driver"
	"github.com/vitaLemoTea/secondstepweb/internal/form"
	"github.com/vitaLemoTea/secondstepweb/internal/helpers"
	"github.com/vitaLemoTea/secondstepweb/internal/model"
	"github.com/vitaLemoTea/secondstepweb/internal/render"
	"github.com/vitaLemoTea/secondstepweb/internal/repository"
	"github.com/vitaLemoTea/secondstepweb/internal/repository/dbrepo"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var Repo *Repository

type Repository struct {
	CF *config.Config
	DB repository.DatabaseRepo
}

func NewHandler(r *Repository) {
	Repo = r
}

func NewRepo(cf *config.Config, db *driver.DB) *Repository {
	return &Repository{
		CF: cf,
		DB: dbrepo.NewPostgresRepo(cf, db.SQL),
	}
}

func NewTestRepo(cf *config.Config) *Repository {
	return &Repository{
		CF: cf,
		DB: dbrepo.NewTestRepo(cf),
	}
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
	err := render.Template(w, "home.page.html", r, &d)
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
	err := render.Template(w, "about.page.html", r, &d)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func (repo *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	data := model.TemplateData{}
	err := render.Template(w, "general.page.html", r, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (repo *Repository) Major(w http.ResponseWriter, r *http.Request) {
	data := model.TemplateData{}
	err := render.Template(w, "major.page.html", r, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (repo *Repository) SearchRooms(w http.ResponseWriter, r *http.Request) {
	data := model.TemplateData{}
	err := render.Template(w, "reservation.page.html", r, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (repo *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	data := model.TemplateData{}
	err := render.Template(w, "reservation.page.html", r, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (repo *Repository) ReservationPage(w http.ResponseWriter, r *http.Request) {
	//id, err := strconv.Atoi(chi.URLParam(r, "id"))

	explored := strings.Split(r.RequestURI, "/")
	id, err := strconv.Atoi(explored[2])

	if err != nil {
		log.Println(err)
		log.Println(r.RequestURI)
		log.Println(r.URL)
		helpers.ServerError(w, err)
		return
	}

	res, ok := repo.CF.Session.Get(r.Context(), "reservation").(model.Reservation)
	if !ok {
		helpers.ServerError(w, errors.New("get reservation from session error"))
		return
	}

	room, err := repo.DB.GetRoomByID(id)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	res.RoomID = id
	res.Room = room
	repo.CF.Session.Put(r.Context(), "reservation", res)
	http.Redirect(w, r, "/RenderResPage", http.StatusTemporaryRedirect)
}

func (repo *Repository) RenderReservationPage(w http.ResponseWriter, r *http.Request) {
	data := model.TemplateData{Data: map[string]interface{}{}}
	res, ok := repo.CF.Session.Get(r.Context(), "reservation").(model.Reservation)
	if !ok {
		helpers.ServerError(w, errors.New("get reservation from session error"))
		return
	}
	f := form.New(nil)
	log.Println(res.RoomID)
	data.Data["reservation"] = res
	data.Data["room_id"] = res.RoomID
	data.Form = f
	render.Template(w, "book.page.html", r, &data)

}

func (repo *Repository) HandleReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	data := model.TemplateData{
		Data: map[string]interface{}{},
	}

	res, ok := repo.CF.Session.Get(r.Context(), "reservation").(model.Reservation)
	if !ok {
		helpers.ServerError(w, errors.New("get reservation from session error"))
		return
	}
	f := form.New(r.PostForm)
	res.FirstName = r.Form.Get("first_name")
	res.LastName = r.Form.Get("last_name")
	res.Email = r.Form.Get("email")
	res.Phone = r.Form.Get("phone")

	f.CheckNotnull("first_name", "last_name", "email", "phone", "room_id")
	f.CheckEmail("email")
	data.Data["reservation"] = res
	data.Form = f
	if !f.Valid() {
		err := render.Template(w, "book.page.html", r, &data)
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
	availability, err := repo.DB.SearchAvailabilityByDates(res.RoomID, res.StartDate, res.EndDate)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	if !availability {
		repo.CF.Session.Put(r.Context(), "error", "your chosened date not available")
		err := render.Template(w, "book.page.html", r, &data)
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
	//insert into db
	id, err := repo.DB.InsertReservation(res)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	roomRestriction := model.RoomRestriction{
		DBModel: model.DBModel{
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
		},
		ReservationID: id,
		RestrictionID: 1,
		StartDate:     res.StartDate,
		EndDate:       res.EndDate,
		RoomID:        res.RoomID,
	}
	err = repo.DB.InsertRoomRestriction(roomRestriction)
	repo.CF.MailSender.SendMail([]string{"jia@163.com"}, "make reservation successed")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	repo.CF.Session.Put(r.Context(), "reservation", res)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)

	return
}

type JsonRes struct {
	Message   string `json:"message,omitempty"`
	OK        bool   `json:"ok,omitempty"`
	RoomID    string `json:"roomID"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

func (repo *Repository) AvailabilityJson(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	tFormat := "2006-01-02"
	startDate := r.Form.Get("start_date")
	endDate := r.Form.Get("end_date")
	roomID, err := strconv.Atoi(r.Form.Get("room_id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	sd, err := time.Parse(tFormat, startDate)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	ed, err := time.Parse(tFormat, endDate)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	avail, err := repo.DB.SearchAvailabilityByDates(roomID, sd, ed)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	var t JsonRes
	if !avail {
		t.Message = "chosen date not availability"
		t.OK = avail
	} else {
		t = JsonRes{
			Message:   "you are so wonderful",
			OK:        avail,
			RoomID:    r.Form.Get("room_id"),
			StartDate: sd.Format("2006-01-02"),
			EndDate:   ed.Format("2006-01-02"),
		}

	}

	bs, err := json.MarshalIndent(t, "", " ")
	log.Println(string(bs))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bs)
}

func (repo *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	tFormat := "2006-01-02"
	startDate := r.Form.Get("start_date")
	endDate := r.Form.Get("end_date")
	sd, err := time.Parse(tFormat, startDate)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	ed, err := time.Parse(tFormat, endDate)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	rooms, err := repo.DB.SerachAvailabilityAllRooms(sd, ed)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	if len(rooms) == 0 {
		repo.CF.Session.Put(r.Context(), "error", "no availability rooms ,please chose other dates!")
		http.Redirect(w, r, "/SearchRooms", http.StatusSeeOther)
		return
	}
	res := model.Reservation{
		StartDate: sd,
		EndDate:   ed,
	}
	repo.CF.Session.Put(r.Context(), "reservation", res)
	data := make(map[string]interface{})
	data["rooms"] = rooms
	render.Template(w, "searchResult.page.html", r, &model.TemplateData{Data: data})
	return
}

func (repo *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := repo.CF.Session.Get(r.Context(), "reservation").(model.Reservation)

	if !ok {
		repo.CF.ErrorLog.Println("this key can not find in the session")
		repo.CF.Session.Put(r.Context(), "error", "can not find session ")
		http.Redirect(w, r, "/book", http.StatusTemporaryRedirect)
		return
	}
	repo.CF.Session.Remove(r.Context(), "reservation")
	data := model.TemplateData{Data: map[string]interface{}{"reservation": reservation}}
	err := render.Template(w, "reservation-summary.page.html", r, &data)
	if err != nil {
		log.Println(err)
	}
}

func (repo Repository) BookRoom(w http.ResponseWriter, r *http.Request) {
	tFormat := "2006-01-02"
	var res model.Reservation
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	//or we can use chi 's helper function
	//chi.URLParam(r,"id")
	id, err := strconv.Atoi(r.Form.Get("id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	std, err := time.Parse(tFormat, r.Form.Get("std"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	ed, err := time.Parse(tFormat, r.Form.Get("ed"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	room, err := repo.DB.GetRoomByID(id)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	res.RoomID = id
	res.StartDate = std
	res.EndDate = ed
	log.Println("sssss:", res)
	res.Room = room
	repo.CF.Session.Put(r.Context(), "reservation", res)
	http.Redirect(w, r, "/RenderResPage", http.StatusSeeOther)
}
