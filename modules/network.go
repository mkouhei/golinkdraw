/*
 Copyright (c) 2014 Kouhei Maeda <mkouhei@palmtb.net>

 This software is release under the Expat License.
*/
package modules

import (
	"strconv"

	"github.com/ajstarks/svgo"
)

func (canv Canvas) Network() *svg.SVG {
	canvas := svg.New(canv.W)
	canvas.Start(canv.Width, canv.Height)

	d := 5 // diameter
	linestyle := "fill:none;stroke:black"
	textstyle := "font-size:10px"
	circlestyle := "fill:gray"

	nd := 100
	xpoly := make([]int, nd)
	ypoly := make([]int, nd)
	xpoly[0] = canv.Width / 2
	ypoly[0] = canv.Height / 2
	delta := 5
	canvas.Circle(xpoly[0], ypoly[0], d, circlestyle)
	for i := 1; i < nd; i++ {
		if i%4 == 1 {
			xpoly[i] = xpoly[i-1] + delta*i
			ypoly[i] = ypoly[i-1] + delta
		} else if i%4 == 2 {
			xpoly[i] = xpoly[i-1] - delta
			ypoly[i] = ypoly[i-1] + delta*i
		} else if i%4 == 3 {
			xpoly[i] = xpoly[i-1] - delta*i
			ypoly[i] = ypoly[i-1] - delta
		} else if i%4 == 0 {
			xpoly[i] = xpoly[i-1] + delta
			ypoly[i] = ypoly[i-1] - delta*i
		}
		canvas.Circle(xpoly[i], ypoly[i], d, circlestyle)
		canvas.Text(xpoly[i]+d*2, ypoly[i]-d*2, "("+strconv.Itoa(xpoly[i])+","+strconv.Itoa(ypoly[i])+")", textstyle)
		if i > 0 {
			canvas.Line(xpoly[i-1], ypoly[i-1], xpoly[i], ypoly[i], linestyle)
		}
	}

	canvas.End()
	return canvas
}
