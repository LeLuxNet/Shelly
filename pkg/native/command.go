package native

import (
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"io"
)

type Native struct{}

func (Native) Run(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer, session *sessions.Session) error {
	return Exec(args, stdin, stdout, stderr)
}

type NoCmd struct {
	Message string
}

func (NoCmd) Code() int {
	return 2
}

func (NoCmd) Error() string {
	return "No such cmd"
}
