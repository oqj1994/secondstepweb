package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMyMiddleWare(t *testing.T) {
	var myH myHandler
	h := MyMiddleWare(&myH)
	switch v := h.(type) {
	case http.Handler:
	//do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.Handler,but type is %T", v))

	}
}

func TestCsrf(t *testing.T) {
	var myH myHandler
	h := Csrf(&myH)
	switch v := h.(type) {
	case http.Handler:
	default:
		t.Error(fmt.Sprintf("type is not http.Handler,but is %T", v))
	}
}

func TestScs(t *testing.T) {
	var myH myHandler
	h := SessionLoad(&myH)
	switch v := h.(type) {
	case http.Handler:
	default:
		t.Error(fmt.Sprintf("type is not http.Handler,but is %T", v))
	}
}
