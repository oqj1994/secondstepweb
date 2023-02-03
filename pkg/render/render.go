package render

import (
	"errors"
	"github.com/vitaLemoTea/secondstepweb/pkg/config"
	"github.com/vitaLemoTea/secondstepweb/pkg/model"
	"html/template"
	"net/http"
	"path/filepath"
)

var cf *config.Config

func NewRender(config *config.Config) {
	cf = config
}

func Render(w http.ResponseWriter, tmpl string, data *model.TemplateData) error {
	var tpl *template.Template
	if cf.UseCache {
		_, ok := cf.TC[tmpl]
		if !ok {
			return errors.New("get tmpl cache failed")
		}
		tpl = cf.TC[tmpl]
	} else {
		tpls, err := GetTc()
		if err != nil {
			return err
		}
		tpl = tpls[tmpl]
	}

	err := tpl.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}

func GetTc() (map[string]*template.Template, error) {
	tc := make(map[string]*template.Template)
	ps, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return tc, err
	}
	for _, p := range ps {
		name := filepath.Base(p)
		tpl, _ := template.New(name).ParseFiles(p)
		ls, err := filepath.Glob("./templates/*.layout.html")
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
