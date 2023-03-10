package render

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/justinas/nosurf"
	"github.com/vitaLemoTea/secondstepweb/internal/config"
	"github.com/vitaLemoTea/secondstepweb/internal/model"
	"html/template"
	"net/http"
	"path/filepath"
)

var cf *config.Config
var pathToTemplate = "./templates"
var functions template.FuncMap

func NewRenderer(config *config.Config) {
	cf = config
}

func AddDefault(data *model.TemplateData, r *http.Request) *model.TemplateData {
	data.Flash = cf.Session.PopString(r.Context(), "flash")
	data.Warning = cf.Session.PopString(r.Context(), "warning")
	data.Error = cf.Session.PopString(r.Context(), "error")
	data.CSRFToken = nosurf.Token(r)
	return data
}

func Template(w http.ResponseWriter, tmpl string, r *http.Request, data *model.TemplateData) error {
	var tpl *template.Template
	if cf.UseCache {
		_, ok := cf.TC[tmpl]
		if !ok {
			return errors.New("get tmpl cache failed")
		}
		tpl = cf.TC[tmpl]
	} else {
		tpls, err := GetTemplateCache()
		if err != nil {
			return err
		}
		tpl = tpls[tmpl]
	}
	data = AddDefault(data, r)
	buf := new(bytes.Buffer)

	err := tpl.Execute(buf, data)
	if err != nil {
		return err
	}
	_, err = w.Write(buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func GetTemplateCache() (map[string]*template.Template, error) {
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
