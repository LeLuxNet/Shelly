package native

import (
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"os/exec"
	"strings"
)

func execProgram(cmd string, args []string, std sessions.Std, dir string) error {
	exe := exec.Command(cmd, args...)
	exe.Stdin = std.In
	exe.Stdout = std.Out
	exe.Stderr = std.Err
	exe.Dir = dir
	return exe.Run()
}

func join(args []string) string {
	var parts []string
	for _, part := range args {
		if strings.Contains(part, " ") {
			parts = append(parts, "\""+part+"\"")
		} else {
			parts = append(parts, part)
		}
	}
	return strings.Join(parts, " ")
}
