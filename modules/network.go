/*
 Copyright (c) 2014 Kouhei Maeda <mkouhei@palmtb.net>

 This software is release under the Expat License.
*/
package modules

import (
	"github.com/ajstarks/svgo"
	"math/rand"
)

func (canv Canvas) Network() *svg.SVG {
	canvas := svg.New(canv.W)
	canvas.Start(canv.Width, canv.Height)

	//x := canv.Width / 20
	//y := canv.Height / 20
	d := 5 // diameter
	//span := 50
	linestyle := "fill:none;stroke:black"
	circlestyle := "fill:gray"

	nd := 100
	xpoly := make([]int, nd)
	ypoly := make([]int, nd)
	for i := 0; i < nd; i++ {
		xpoly[i] = rand.Intn(canv.Width)
		ypoly[i] = rand.Intn(canv.Height)
		canvas.Circle(xpoly[i], ypoly[i], d, circlestyle)
	}
	/*
		xpoly[0] = x
		ypoly[0] = y
		xpoly[1] = x + span*2
		ypoly[1] = y + span*3
		xpoly[2] = x + span*4
		ypoly[2] = y + span*5
		canvas.Circle(xpoly[0], ypoly[0], d, circlestyle)
		canvas.Circle(xpoly[1], ypoly[1], d, circlestyle)
		canvas.Circle(xpoly[2], ypoly[2], d, circlestyle)
	*/
	canvas.Polyline(xpoly, ypoly, linestyle)

	canvas.End()
	return canvas
}
