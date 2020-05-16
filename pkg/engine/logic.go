package engine

import (
	"github.com/LeLuxNet/Shelly/pkg/command"
	"github.com/LeLuxNet/Shelly/pkg/output"
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"regexp"
	"strings"
)

func MultiLineInput(text string, session *sessions.Session) {
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

func singleLineInput(line string, session *sessions.Session) {
	sessions.AddHistory(line)
	ands := strings.Split(line, " && ")
	for _, and := range ands {
		code := singleCommandInput(and, session)
		if code != 0 {
			break
		}
	}
}

func singleCommandInput(cmd string, session *sessions.Session) int {
	regex := regexp.MustCompile(`\s+`)
	args := regex.Split(strings.TrimSpace(cmd), -1)
	newCmd, exe := command.GetRegisteredNative(args[0])
	args[0] = newCmd
	err := exe.Run(args, session.In, session.Out, session.Err, session)
	if err != nil {
		// TODO: Add error code
		output.SendNl(output.Color(err.Error(), output.COLOR_F_RED), session.Err)
		return 1
	}
	return 0
}
