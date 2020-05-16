package command

import (
	"github.com/LeLuxNet/Shelly/pkg/native"
	"strings"
)

const (
	NativePrefix = "."
)

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

func GetRegisteredNative(listen string) (string, Cmd) {
	if strings.HasPrefix(listen, NativePrefix) {
		return strings.TrimPrefix(listen, NativePrefix), native.Native{}
	}
	cmd := GetRegistered(listen)
	if cmd == nil {
		cmd = native.Native{}
	}
	return listen, cmd
}
