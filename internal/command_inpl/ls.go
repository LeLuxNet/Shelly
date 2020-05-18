package command_inpl

import (
	"github.com/LeLuxNet/Shelly/pkg/output"
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"os"
	"strings"
)

type Ls struct{}

func (Ls) Run(args []string, std sessions.Std, session *sessions.Session) error {
	files, err := session.WorkingDir.ListDir(false)
	if err != nil {
		return err
	}

	for _, file := range files {
		contentType := ""
		path, err := session.WorkingDir.GetRelativePath(file.Name(), true)
		if err != nil {
			return err
		}
		if path.ExpectDir(false) == nil {
			contentType, err = path.ContentType()
			if err != nil {
				return err
			}
		}
		var result string
		if path.ExpectDir(true) == nil {
			result = output.GetColor(output.COLOR_F_BLUE)
		} else if file.Mode() == os.ModeSymlink {
			result = output.GetColor(output.COLOR_F_CYAN)
		} else if file.Mode() == os.ModeDevice {
			result = output.GetColor(output.COLOR_FB_YELLOW, output.COLOR_B_BLACK)
		} else if strings.HasPrefix(contentType, "image/") {
			result = output.GetColor(output.COLOR_FB_MAGENTA)
		} else if file.Mode()&0111 != 0 {
			result = output.GetColor(output.COLOR_FB_GREEN)
		} else {
			result = output.GetColor(output.COLOR_RESET)
		}
		output.SendNl(result+file.Name(), std.Out)
	}
	return nil
}
