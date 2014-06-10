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

	d := 5 // diameter
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
	canvas.Polyline(xpoly, ypoly, linestyle)

	canvas.End()
	return canvas
}
