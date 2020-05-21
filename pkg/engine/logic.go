package engine

import (
	"github.com/LeLuxNet/Shelly/pkg/command"
	"github.com/LeLuxNet/Shelly/pkg/output"
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"regexp"
	"strings"
)

func MultiLineInput(text string, std sessions.Std, session *sessions.Session, history bool) {
	text = strings.TrimSpace(strings.ReplaceAll(text, "\\\n", ""))
	if text == "" {
		return
	}
	regex := regexp.MustCompile(`( *(?:;|\n|\r\n) *)`)
	lines := regex.Split(text, -1)
	for _, line := range lines {
		singleLineInput(line, std, session, history)
	}
}

func singleLineInput(line string, std sessions.Std, session *sessions.Session, history bool) {
	if history {
		sessions.AddHistory(line)
	}
	ands := strings.Split(line, " && ")
	for _, and := range ands {
		code := singleCommandInput(and, std, session)
		if code != 0 {
			break
		}
	}
}

func singleCommandInput(cmd string, std sessions.Std, session *sessions.Session) int {
	args := split(cmd, ' ', '"')
	newCmd, exe := command.GetRegisteredNative(args[0])
	args[0] = newCmd
	err := exe.Run(args, std, session)
	if err != nil {
		// TODO: Add error code
		output.SendNl(output.Color(err.Error(), output.COLOR_F_RED), std.Err)
		return 1
	}
	return 0
}

func split(input string, sep rune, stick rune) []string {
	var result []string
	part := ""
	stickActive := false
	for _, char := range input {
		if stickActive {
			if char == stick {
				stickActive = false
			} else {
				part += string(char)
			}
		} else {
			if char == stick && part == "" {
				stickActive = true
			} else if char == sep && part != "" {
				result = append(result, part)
				part = ""
			} else {
				part += string(char)
			}
		}
	}
	return append(result, part)
}
