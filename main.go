/*
 github.com/mkouhei/golinkdraw/main.go

 Copyright (c) 2014 Kouhei Maeda <mkouhei@palmtb.net>

 This software is release under the Expat License.
*/
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
	go h.run()
	http.Handle("/", http.HandlerFunc(serveHome))
	http.Handle("/ws", http.HandlerFunc(serveWs))
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	if err := http.ListenAndServe(s.Addr, nil); err != nil {
		log.Fatal("ListenAndServe: %s", err)
	}
}
