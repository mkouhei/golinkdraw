package main

import (
	"log"
	"net/http"
	"text/template"
)


var (
	s = &http.Server{
		Addr: ":8080",
	}
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
		StringSVG(),
	}
	homeTempl.Execute(w, &v)
}

func main() {
	http.Handle("/", http.HandlerFunc(serveHome))
	http.Handle("/ws", http.HandlerFunc(serveWs))
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	if err := http.ListenAndServe(s.Addr, nil); err != nil {
		log.Fatal("ListenAndServe: %s", err)
	}
}
