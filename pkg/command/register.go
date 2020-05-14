package command

import "github.com/LeLuxNet/Shelly/pkg/native"

var registeredCmds = make(map[string]Cmd)

func Register(listen string, cmd Cmd) bool {
	if _, ok := registeredCmds[listen]; ok {
		return false
	}
	registeredCmds[listen] = cmd
	return true
}

func GetRegistered(listen string) Cmd {
	return registeredCmds[listen]
}

var NativeCmd = native.Native{}