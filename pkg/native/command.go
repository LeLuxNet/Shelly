package native

import (
	"github.com/LeLuxNet/Shelly/pkg/sessions"
)

type Native struct{}

func (Native) Run(args []string, std sessions.Std, session *sessions.Session) error {
	return Exec(args, std, session.WorkingDir.General, session)
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
