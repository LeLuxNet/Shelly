package engine

import (
	"github.com/LeLuxNet/Shelly/pkg/command"
	"github.com/LeLuxNet/Shelly/pkg/output"
	"github.com/LeLuxNet/Shelly/pkg/session"
	"regexp"
	"strings"
)

func MultiLineInput(text string, session *session.Session) {
	text = strings.TrimSpace(strings.ReplaceAll(text, "\\\n", ""))
	if text == "" {
		return
	}
	regex := regexp.MustCompile(`( *[;\n] *)`)
	lines := regex.Split(text, -1)
	for _, line := range lines {
		singleLineInput(line, session)
	}
}

func singleLineInput(line string, s *session.Session) {
	session.AddHistory(line)
	ands := strings.Split(line, " && ")
	for _, and := range ands {
		code := singleCommandInput(and, s)
		if code != 0 {
			break
		}
	}
}

func singleCommandInput(cmd string, session *session.Session) int {
	regex := regexp.MustCompile(`\s+`)
	args := regex.Split(strings.TrimSpace(cmd), -1)
	exec := command.GetRegistered(args[0])
	if exec == nil {
		exec = command.NativeCmd
	}
	err := exec.Run(args, session.In, session.Out, session.Err, session)
	if err != nil {
		// TODO: Add error code
		output.SendNl("Error: "+err.Error(), session.Err)
		return 1
	}
	return 0
}
