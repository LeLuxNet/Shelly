package command_inpl

import (
	"github.com/LeLuxNet/Shelly/pkg/output"
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"io"
	"os"
)

type Ls struct{}

func (Ls) Run(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer, session *sessions.Session) error {
	files, err := session.WorkingDir.ListDir(false)
	if err != nil {
		return err
	}

	result := ""
	for _, file := range files {
		if file.IsDir() {
			result += output.GetColor(output.COLOR_F_BLUE)
		} else if file.Mode() == os.ModeSymlink {
			result += output.GetColor(output.COLOR_F_CYAN)
		} else if file.Mode() == os.ModeDevice {
			result += output.GetColor(output.COLOR_FB_YELLOW, output.COLOR_B_BLACK)
		} else if file.Mode()&0111 != 0 {
			result += output.GetColor(output.COLOR_FB_GREEN)
		} else {
			result += output.GetColor(output.COLOR_RESET)
		}
		result += file.Name() + "\n"
	}
	output.Send(result, stdout)
	return nil
}
