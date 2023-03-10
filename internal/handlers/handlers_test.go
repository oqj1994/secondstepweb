package handlers

import (
	"context"
	"github.com/vitaLemoTea/secondstepweb/internal/model"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTest = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{
		name:               "home",
		url:                "/",
		method:             "Get",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "about",
		url:                "/about",
		method:             "Get",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "generals",
		url:                "/generals",
		method:             "Get",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "majors",
		url:                "/majors",
		method:             "Get",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "make-reservation",
		url:                "/make-reservation",
		method:             "Get",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	}, {
		name:               "contact",
		url:                "/contact",
		method:             "Get",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
}

func TestGetMethod(t *testing.T) {
	route := GetRoutes()
	ts := httptest.NewServer(route)
	defer ts.Close()
	for _, e := range theTest {
		if e.method == "Get" {
			//do the get request
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s,expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			pData := url.Values{}
			for _, v := range e.params {
				pData.Set(v.key, v.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, pData)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s,expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}

func TestMakeReservation(t *testing.T) {

	req, err := http.NewRequest("GET", "/make-reservation/1", nil)
	req.RequestURI = "/make-reservation/1"
	if err != nil {
		log.Fatal("create request error")
	}
	ctx := GetContext(req)
	reservation := model.Reservation{
		ID:     0,
		RoomID: 0,
		Room:   model.Room{},
	}
	Repo.CF.Session.Put(ctx, "reservation", reservation)
	req = req.WithContext(ctx)
	response := httptest.ResponseRecorder{}
	handler := http.HandlerFunc(Repo.ReservationPage)
	handler.ServeHTTP(&response, req)
	if response.Code != http.StatusTemporaryRedirect {
		t.Error("can not get expect code")
	}

}

func TestRepository_RenderReservationPage(t *testing.T) {
	req := httptest.NewRequest("GET", "/RenderResPage", nil)
	ctx := GetContext(req)
	reservation := model.Reservation{
		ID:     0,
		RoomID: 0,
		Room:   model.Room{},
	}
	Repo.CF.Session.Put(ctx, "reservation", reservation)
	req = req.WithContext(ctx)
	responer := httptest.ResponseRecorder{}
	handler := http.HandlerFunc(Repo.RenderReservationPage)
	handler.ServeHTTP(&responer, req)
	if responer.Code != http.StatusOK {
		t.Errorf("expect to get http code 200,but get %d", responer.Code)
	}
}

func GetContext(r *http.Request) context.Context {

	ctx, err := appConfig.Session.Load(r.Context(), r.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}
