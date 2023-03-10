package render

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/vitaLemoTea/secondstepweb/internal/config"
	"github.com/vitaLemoTea/secondstepweb/internal/model"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

var session *scs.SessionManager
var appConfig config.Config

func TestMain(t *testing.M) {
	gob.Register(model.Reservation{})
	appConfig.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.HttpOnly = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.InProduction

	appConfig.Session = session
	appConfig.UseCache = true
	info := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	appConfig.InfoLog = info
	errorLog := log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
	appConfig.ErrorLog = errorLog
	cf = &appConfig
	os.Exit(t.Run())
}

type myResponseWriter struct {
}

func (m *myResponseWriter) Header() http.Header {
	var h http.Header
	return h
}

func (m *myResponseWriter) Write(bytes []byte) (int, error) {
	length := len(bytes)
	return length, nil
}

func (m *myResponseWriter) WriteHeader(statusCode int) {

}
