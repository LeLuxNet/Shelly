package main

import (
	"os"
	"path"
	"strings"
)

type Path struct {
	General string
}

func NewPath(path string) *Path {
	return &Path{General: strings.ReplaceAll(path, "\\", "/")}
}

func (p *Path) ChDir(relative string) CmdCrashError {
	target := path.Join(p.General, relative)
	if _, err := os.Stat(target); os.IsNotExist(err) {
		return PathError{NotExists}
	}
	p.General = target
	return nil
}

func (p *Path) CurrentDir() string {
	dirs := strings.Split(p.General, "/")
	return dirs[len(dirs)-1]
}
