package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"testing"
)

func TestNewRouter(t *testing.T) {

	h := NewRouter()
	switch v := h.(type) {
	case chi.Router:
	default:
		t.Error(fmt.Sprintf("type is not chi.Router,type is %T", v))
	}
}
