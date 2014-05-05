package modules

import (
	"io"
	"time"

	"github.com/ajstarks/svgo"
)

func now() string {
	t := time.Now()
	return t.Format(time.RFC3339)
}

func SimpleCircle(w io.Writer) *svg.SVG {
	canvas := svg.New(w)
	width := 400
	height := 400
	canvas.Start(width, height)
	canvas.Circle(width/2, height/2, 200)
	canvas.Text(width/2, height/2, now(),
		"text-anchor:middle; font-size: 16px; fill: white")
	canvas.End()
	return canvas
}
