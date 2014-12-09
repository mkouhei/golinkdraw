package main

/*
 github.com/mkouhei/golinkdraw/main.go

 Copyright (c) 2014 Kouhei Maeda <mkouhei@palmtb.net>

 This software is release under the Expat License.
*/

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	s = &http.Server{
		Addr: ":8080",
	}
)

var version string
var showVersion = flag.Bool("version", false, "show_version")

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Printf("version: %s\n", version)
		return
	}
	go h.run()
	http.Handle("/", http.HandlerFunc(serveHome))
	http.Handle("/ws", http.HandlerFunc(serveWs))
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	if err := http.ListenAndServe(s.Addr, nil); err != nil {
		log.Fatalf("ListenAndServe: %s", err)
	}
}
