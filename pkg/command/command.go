package command

import (
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"io"
)

type Cmd interface {
	Run(args []string, in io.Reader, out io.Writer, err io.Writer, session *sessions.Session) error
}
