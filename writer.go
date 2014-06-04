/*
 github.com/mkouhei/golinkdraw/writer.go

 Copyright (c) 2014 Kouhei Maeda <mkouhei@palmtb.net>

 This software is release under the Expat License.
*/
package main

import (
	"bytes"
	"github.com/mkouhei/golinkdraw/modules"
	"io"
	"os"
)

//func StringSVG(f func(io.Writer) *svg.SVG) string {
func StringSVG(width int, height int) string {
	r, w, _ := os.Pipe()
	os.Stdout = w

	// rendering SVG
	canv := modules.Canvas{width, height, os.Stdout}
	//canv.SimpleCircle()
	canv.Richter()

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
