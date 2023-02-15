package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
)

type Config struct {
	TC           map[string]*template.Template
	UseCache     bool
	Session      *scs.SessionManager
	InProduction bool
}
