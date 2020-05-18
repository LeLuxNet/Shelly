package path

import (
	"io"
	"net/http"
	"os"
)

func (p Path) ContentType() (string, error) {
	err := p.ExpectDir(false)
	if err != nil {
		return "", err
	}
	file, err := os.Open(p.General)
	if err != nil {
		return "", err
	}
	defer file.Close()
	buf := make([]byte, 512)
	_, err = file.Read(buf)
	if err != nil && err != io.EOF {
		return "", err
	}
	return http.DetectContentType(buf), nil
}
