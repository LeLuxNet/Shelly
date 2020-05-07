package command

import (
	"github.com/LeLuxNet/Shelly/pkg/models"
	"io"
)

type Cmd interface {
	Run(args []string, in io.Reader, out io.Writer, err io.Writer, session *models.Session) error
}
