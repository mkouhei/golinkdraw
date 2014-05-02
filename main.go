package main

import (
	"log"
	"net/http"
)

var (
	s = &http.Server{
		Addr: ":8080",
	}
)

func main() {
	http.Handle("/", http.HandlerFunc(serveHome))
	http.Handle("/ws", http.HandlerFunc(serveWs))
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	if err := http.ListenAndServe(s.Addr, nil); err != nil {
		log.Fatal("ListenAndServe: %s", err)
	}
}
