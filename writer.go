package main

/*
 github.com/mkouhei/golinkdraw/writer.go

 Copyright (c) 2014 Kouhei Maeda <mkouhei@palmtb.net>

 This software is release under the Expat License.
*/

import (
	"bytes"

	"github.com/mkouhei/golinkdraw/modules"
)

func stringSVG(width int, height int) string {
	buf := &bytes.Buffer{}
	// rendering SVG
	canv := modules.Canvas{width, height, buf}
	//canv.SimpleCircle()
	//canv.Richter()
	canv.Network()

	return string(buf.Bytes())
}
