// richter -- inspired by Gerhard Richter's 256 colors, 1974
// https://github.com/ajstarks/svgo/richter/richter.go

package modules

import (
	"math/rand"
	"time"

	"github.com/ajstarks/svgo"
)

func (canv Canvas) Richter() *svg.SVG {
	canvas := svg.New(canv.W)
	rand.Seed(int64(time.Now().Nanosecond()) % 1e9)
	canvas.Start(canv.Width, canv.Height)
	canvas.Title("Richter")
	canvas.Rect(0, 0, canv.Width, canv.Height, "fill:white")
	rw := 32
	rh := 18
	margin := 5
	for i, x := 0, 20; i < 16; i++ {
		x += (rw + margin)
		for j, y := 0, 20; j < 16; j++ {
			canvas.Rect(x, y, rw, rh, canvas.RGB(rand.Intn(255), rand.Intn(255), rand.Intn(255)))
			y += (rh + margin)
		}
	}
	canvas.End()
	return canvas
}
