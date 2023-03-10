package handlers

import (
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"
	"github.com/vitaLemoTea/secondstepweb/internal/config"
	"github.com/vitaLemoTea/secondstepweb/internal/model"
	"github.com/vitaLemoTea/secondstepweb/internal/render"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"
)

var appConfig config.Config
var s *scs.SessionManager
var m *chi.Mux
var nRepo *Repository

func TestMain(t *testing.M) {
	gob.Register(model.Reservation{})
	appConfig.InProduction = false
	s = scs.New()
	s.Lifetime = 24 * time.Hour
	s.Cookie.Persist = true
	s.Cookie.HttpOnly = true
	s.Cookie.SameSite = http.SameSiteLaxMode
	s.Cookie.Secure = appConfig.InProduction
	appConfig.Session = s
	tc, _ := GetTestTc()

	appConfig.TC = tc
	appConfig.UseCache = true
	info := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	appConfig.InfoLog = info
	errorLog := log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
	appConfig.ErrorLog = errorLog
	//db, err := driver.ConnectSql("host=localhost port=5432 user=root password=123456 dbname=bookings sslmode=disable")
	//if err != nil {
	//	log.Fatalln(err)
	//}

	render.NewRenderer(&appConfig)
	nRepo = NewTestRepo(&appConfig)
	NewHandler(nRepo)
	m = chi.NewRouter()
	m.Use(middleware.Logger)
	m.Use(MyMiddleWare)
	//mux.Use(Csrf)
	m.Use(SessionLoad)
	fileServer := http.FileServer(http.Dir("./statics/"))
	m.Handle("/statics/*", http.StripPrefix("/statics/", fileServer))
	m.Get("/", Repo.Home)
	m.Get("/about", Repo.About)
	m.Get("/generals", Repo.Generals)
	m.Get("/majors", Repo.Major)
	m.Get("/make-reservation", Repo.SearchRooms)
	m.Get("/contact", Repo.Contact)
	m.Get("/book", Repo.ReservationPage)
	m.Post("/book", Repo.HandleReservation)
	m.Post("/search", Repo.AvailabilityJson)
	m.Get("/reservation-summary", Repo.ReservationSummary)
	os.Exit(t.Run())
}

func GetRoutes() chi.Router {
	gob.Register(model.Reservation{})
	appConfig.InProduction = false
	s = scs.New()
	s.Lifetime = 24 * time.Hour
	s.Cookie.Persist = true
	s.Cookie.HttpOnly = true
	s.Cookie.SameSite = http.SameSiteLaxMode
	s.Cookie.Secure = appConfig.InProduction
	appConfig.Session = s
	tc, _ := GetTestTc()

	appConfig.TC = tc
	appConfig.UseCache = true
	info := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	appConfig.InfoLog = info
	errorLog := log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
	appConfig.ErrorLog = errorLog
	//db, err := driver.ConnectSql("host=localhost port=5432 user=root password=123456 dbname=bookings sslmode=disable")
	//if err != nil {
	//	log.Fatalln(err)
	//}

	render.NewRenderer(&appConfig)
	nRepo = NewTestRepo(&appConfig)
	NewHandler(nRepo)
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(MyMiddleWare)
	//mux.Use(Csrf)
	mux.Use(SessionLoad)
	fileServer := http.FileServer(http.Dir("./statics/"))
	mux.Handle("/statics/*", http.StripPrefix("/statics/", fileServer))
	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/generals", Repo.Generals)
	mux.Get("/majors", Repo.Major)
	mux.Get("/make-reservation", Repo.SearchRooms)
	mux.Get("/contact", Repo.Contact)
	mux.Get("/book", Repo.ReservationPage)
	mux.Post("/book", Repo.HandleReservation)
	mux.Post("/search", Repo.AvailabilityJson)
	mux.Get("/reservation-summary", Repo.ReservationSummary)
	return mux
}

func MyMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ipAdrre := r.RemoteAddr
		fmt.Println(ipAdrre)
		next.ServeHTTP(w, r)

	})
}

var pathToTemplate = "./../../templates"
var functions = template.FuncMap{}

func GetTestTc() (map[string]*template.Template, error) {
	tc := make(map[string]*template.Template)
	ps, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", pathToTemplate))
	if err != nil {
		return tc, err
	}
	for _, p := range ps {
		name := filepath.Base(p)
		tpl, _ := template.New(name).Funcs(functions).ParseFiles(p)
		ls, err := filepath.Glob(fmt.Sprintf("%s/*.layout.html", pathToTemplate))
		if err != nil {
			return tc, err
		}
		tpl, err = tpl.ParseFiles(ls...)
		if err != nil {
			return tc, err
		}
		tc[name] = tpl
	}
	return tc, nil
}

func Csrf(next http.Handler) http.Handler {
	n := nosurf.New(next)
	n.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Secure:   appConfig.InProduction,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})
	return n
}

func SessionLoad(next http.Handler) http.Handler {
	return appConfig.Session.LoadAndSave(next)
}
