/*
 github.com/mkouhei/golinkdraw/writer.go

 Copyright (c) 2014 Kouhei Maeda <mkouhei@palmtb.net>

 This software is release under the Expat License.
*/
package main

import (
	"bytes"
	"io"
	"os"

	"github.com/ajstarks/svgo"
)

func StringSVG(f func(io.Writer) *svg.SVG) string {
	r, w, _ := os.Pipe()
	os.Stdout = w

	// rendering SVG
	f(os.Stdout)

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
