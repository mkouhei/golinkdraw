package main

import (
	"github.com/mkouhei/golinkdraw/modules"
	"net/http"
	"text/template"
)

var (
	homeTempl = template.Must(template.ParseFiles("templates/index.html"))
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var v = struct {
		Host string
		Data string
	}{
		r.Host,
		StringSVG(modules.RenderingSVG),
	}
	homeTempl.Execute(w, &v)
}
