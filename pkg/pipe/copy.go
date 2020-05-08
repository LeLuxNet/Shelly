package pipe

import (
	"io"
)

const bufferSize = 32 * 1024

func CopyMod(dst io.Writer, src io.Reader, mod func([]byte) []byte) error {
	buf := make([]byte, bufferSize)
	for {
		l, errRead := src.Read(buf)
		data := mod(buf[0:l])
		_, errWrite := dst.Write(data)
		if errRead != nil {
			if errRead == io.EOF {
				break
			}
			return errRead
		}
		if errWrite != nil {
			return errWrite
		}
	}
}
