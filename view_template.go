package main

/*
 github.com/mkouhei/golinkdraw/view_template.go

 Copyright (c) 2014 Kouhei Maeda <mkouhei@palmtb.net>

 This software is release under the Expat License.

 This source code is a modification of the source code of
 https://github.com/gorilla/websocket/examples/filewatch/main.go

 The original source code is licensed under the 2-Clause BSD License,
 and copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
*/

import (
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
		stringSVG(defaultWidth, defaultHeight),
	}
	homeTempl.Execute(w, &v)
}
