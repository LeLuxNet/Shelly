package command

import (
	"github.com/LeLuxNet/Shelly/pkg/sessions"
)

type Cmd interface {
	Run(args []string, std sessions.Std, session *sessions.Session) error
}
