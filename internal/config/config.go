package config

import (
	"github.com/alexedwards/scs/v2"
	"github.com/vitaLemoTea/secondstepweb/internal/mail"
	"html/template"
	"log"
)

type Config struct {
	TC           map[string]*template.Template
	UseCache     bool
	Session      *scs.SessionManager
	InProduction bool
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	MailSender   mail.Mailer
}
