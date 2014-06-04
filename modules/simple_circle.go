/*
 github.com/mkouhei/golinkdraw/modules/simple_circle.go

 Copyright (c) 2014 Kouhei Maeda <mkouhei@palmtb.net>

 This software is release under the Expat License.
*/
package modules

import (
	"github.com/ajstarks/svgo"
	"time"
)

func now() string {
	t := time.Now()
	return t.Format(time.RFC3339)
}

func (canv Canvas) SimpleCircle() *svg.SVG {
	canvas := svg.New(canv.W)
	canvas.Start(canv.Width, canv.Height)
	canvas.Circle(canv.Width/2, canv.Height/2, 200)
	canvas.Text(canv.Width/2, canv.Height/2, now(),
		"text-anchor:middle; font-size: 16px; fill: white")
	canvas.End()
	return canvas
}
