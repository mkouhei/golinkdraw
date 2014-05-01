package main

import (
	"bytes"
	"io"
	"os"
)

func StringSVG() string {
	r, w, _ := os.Pipe()
	os.Stdout = w

	// rendering SVG
	RenderingSVG(os.Stdout)

	o := make(chan string)
	go func() {
		buf := new(bytes.Buffer)
		io.Copy(buf, r)
		o <- buf.String()
	}()
	w.Close()
	strSVG := <-o
	return strSVG
}
