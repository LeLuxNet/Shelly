package sessions

import "io"

type Std struct {
	In  io.Reader
	Out io.Writer
	Err io.Writer
}
