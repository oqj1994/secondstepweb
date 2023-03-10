package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

func MyMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ipAdrre := r.RemoteAddr
		fmt.Println(ipAdrre)
		next.ServeHTTP(w, r)

	})
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
