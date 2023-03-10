package render

import (
	"fmt"
	"github.com/vitaLemoTea/secondstepweb/internal/model"
	"net/http"
	"testing"
)

func TestAddDefault(t *testing.T) {
	td := &model.TemplateData{}
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	session.Put(r.Context(), "flash", "I'm testing the AddDefault function")
	result := AddDefault(td, r)
	if result.Flash != "I'm testing the AddDefault function" {
		t.Error("failed")
	}
}

func TestRender(t *testing.T) {
	pathToTemplate = "./../../templates"
	tc, err := GetTemplateCache()
	if err != nil {
		t.Error(err)
	}
	td := &model.TemplateData{}
	appConfig.TC = tc
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	w := myResponseWriter{}
	err = Template(&w, "about.page.html", r, td)
	if err != nil {
		t.Error("error rendering template to browser")
	}
	err = Template(&w, "f.page.html", r, td)
	if err == nil {
		t.Error("should get error but get nil")
	}
}

func TestGetTc(t *testing.T) {
	pathToTemplate = "./../../templates"
	tc, err := GetTemplateCache()
	if err != nil {
		t.Error("get template cache failed")
	}
	if tc == nil {
		t.Error("get template cache failed")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("Get", "/test", nil)
	if err != nil {
		return nil, err
	}
	ctx := r.Context()
	x := r.Header.Get("X-Session")
	fmt.Println("X-Session is ", x)
	ctx, _ = session.Load(ctx, x)
	r = r.WithContext(ctx)
	return r, nil
}
