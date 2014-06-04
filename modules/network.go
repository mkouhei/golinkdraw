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

	/*
		TODO: implement linkdraw algorism.
	*/

	canvas.End()
	return canvas
}
