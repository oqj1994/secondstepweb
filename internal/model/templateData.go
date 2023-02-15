package model

import (
	"github.com/justinas/nosurf"
	"net/http"
)

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float64
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

func AddDefault(data *TemplateData, r *http.Request) *TemplateData {
	data.CSRFToken = nosurf.Token(r)
	return data
}
