package command

import (
	"github.com/LeLuxNet/Shelly/pkg/session"
	"io"
)

type Cmd interface {
	Run(args []string, in io.Reader, out io.Writer, err io.Writer, session *session.Session) error
}
