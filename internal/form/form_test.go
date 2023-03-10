package form

import (
	"net/http"
	"net/url"
	"testing"
)

func TestNew(t *testing.T) {
	values := url.Values{}
	f := New(values)
	if f == nil {
		t.Error("Get new form failed")
	}
}

func TestForm_Valid(t *testing.T) {
	values := url.Values{}
	f := New(values)
	if !f.Valid() {
		t.Error("valid form error")
	}
}

func TestForm_CheckNotnull(t *testing.T) {
	values := url.Values{}
	values.Set("a", "a")
	values.Set("b", "b")
	values.Set("c", "c")
	f := New(values)
	f.CheckNotnull("a", "b", "c")
	if !f.Valid() {
		t.Error("error check field ")
	}
	if f.Errs.Get("a") != "" {
		t.Error("should get no error")
	}
	values2 := url.Values{}
	values2.Set("a", "a")
	values2.Set("b", "b")
	values2.Set("c", "c")
	f2 := New(values2)
	f2.CheckNotnull("a", "b", "c", "d")
	if f2.Valid() {
		t.Error("expect to get not valid but get valid")
	}
	if f2.Errs.Get("d") == "" {
		t.Error("expect to get error field,but get no error")
	}
}

func TestForm_CheckEmail(t *testing.T) {
	values := url.Values{}
	values.Set("email", "oqj@qq.com")
	r, err := http.NewRequest("POST", "/something", nil)
	if err != nil {
		t.Fail()
	}
	r.Form = values
	f := New(values)
	f.CheckEmail("email")
	if !f.Valid() {
		t.Error("expect to pass the check but get unpass")
	}
	r.Form.Set("email", "oqj.com")
	f.CheckEmail("email")
	if f.Valid() {
		t.Error("expect to unpass the check but get pass")
	}
}
