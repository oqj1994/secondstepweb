package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/vitaLemoTea/secondstepweb/pkg/config"
	"github.com/vitaLemoTea/secondstepweb/pkg/handlers"
	"github.com/vitaLemoTea/secondstepweb/pkg/render"
	"log"
	"net/http"
	"time"
)

type App struct {
	c   chan error
	mux chi.Router
}

func (a App) Run() {
	port := ":8080"
	fmt.Println("server start running on port ", port)
	a.c <- http.ListenAndServe(port, a.mux)
}

var appConfig config.Config

func main() {

	appConfig.InProduction = false
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.InProduction
	appConfig.Session = session
	tc, err := render.GetTc()
	if err != nil {
		log.Fatal(err)
	}
	appConfig.TC = tc
	appConfig.UseCache = false

	render.NewRender(&appConfig)
	nRepo := handlers.NewRepo(&appConfig)
	handlers.NewHandler(nRepo)
	mux := NewRouter()
	app := App{
		c:   make(chan error),
		mux: mux,
	}
	go app.Run()
	if err := <-app.c; err != nil {
		fmt.Println("server error:", err)
	}
}
