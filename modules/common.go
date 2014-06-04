package modules

import (
	"io"
)

type Canvas struct {
	Width  int
	Height int
	W      io.Writer
}
