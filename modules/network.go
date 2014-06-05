/*
 Copyright (c) 2014 Kouhei Maeda <mkouhei@palmtb.net>

 This software is release under the Expat License.
*/
package modules

import (
	"github.com/ajstarks/svgo"
)

func (canv Canvas) Network() *svg.SVG {
	canvas := svg.New(canv.W)
	canvas.Start(canv.Width, canv.Height)

	x := canv.Width / 20
	y := canv.Height / 20
	d := 5 // diameter
	span := 50
	linestyle := "fill:none;stroke:black"
	circlestyle := "fill:gray"

	xpoly := make([]int, 3)
	ypoly := make([]int, 3)
	xpoly[0] = x
	ypoly[0] = y
	xpoly[1] = x + span*2
	ypoly[1] = y + span*3
	xpoly[2] = x + span*4
	ypoly[2] = y + span*5

	canvas.Circle(xpoly[0], ypoly[0], d, circlestyle)
	canvas.Circle(xpoly[1], ypoly[1], d, circlestyle)
	canvas.Circle(xpoly[2], ypoly[2], d, circlestyle)
	canvas.Polyline(xpoly, ypoly, linestyle)

	canvas.End()
	return canvas
}
