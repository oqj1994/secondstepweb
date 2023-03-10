package main

import (
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/vitaLemoTea/secondstepweb/internal/config"
	"github.com/vitaLemoTea/secondstepweb/internal/driver"
	"github.com/vitaLemoTea/secondstepweb/internal/handlers"
	"github.com/vitaLemoTea/secondstepweb/internal/helpers"
	"github.com/vitaLemoTea/secondstepweb/internal/mail"
	"github.com/vitaLemoTea/secondstepweb/internal/model"
	"github.com/vitaLemoTea/secondstepweb/internal/render"
	"log"
	"net/http"
	"os"
	"time"
)

type App struct {
	c   chan error
	mux chi.Router
}

func (a App) Run() {
	port := ":8081"
	fmt.Println("server start running on port ", port)
	a.c <- http.ListenAndServe(port, a.mux)
}

var appConfig config.Config
var mux chi.Router
var info *log.Logger
var errorLog *log.Logger

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()
	app := App{
		c:   make(chan error),
		mux: mux,
	}
	go app.Run()
	if err := <-app.c; err != nil {
		fmt.Println("server error:", err)
	}
}

func run() (*driver.DB, error) {
	
	gob.Register(model.Reservation{})
	gob.Register(model.User{})
	gob.Register(model.Room{})
	gob.Register(model.RoomRestriction{})
	gob.Register(model.Restriction{})

	appConfig.InProduction = false

	mailSender := mail.InitMailSender()
	appConfig.MailSender = mailSender
	session := InitSession()
	appConfig.Session = session
	tc, err := render.GetTemplateCache()

	if err != nil {
		return nil, err
	}
	info = log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	appConfig.InfoLog = info
	errorLog = log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
	appConfig.ErrorLog = errorLog
	appConfig.TC = tc
	appConfig.UseCache = false
	log.Println("connect to the database")
	dsn := "host=localhost port=5432 user=root password=123456 dbname=bookings sslmode=disable"
	db, err := driver.ConnectSql(dsn)
	if err != nil {
		log.Fatalln("can not connect to the database")
	}
	log.Println("connected to the database")

	render.NewRenderer(&appConfig)
	nRepo := handlers.NewRepo(&appConfig, db)
	helpers.NewHelpers(&appConfig)
	handlers.NewHandler(nRepo)
	mux = NewRouter()
	return db, nil
}

func InitSession() *scs.SessionManager {
	s := scs.New()
	s.Lifetime = 24 * time.Hour
	s.Cookie.Persist = true
	s.Cookie.HttpOnly = true
	s.Cookie.SameSite = http.SameSiteLaxMode
	s.Cookie.Secure = appConfig.InProduction
	return s
}
