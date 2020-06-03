package path

import (
	"os/user"
	"runtime"
	"strings"
)

type VPath struct {
	*Path
	Visible string
}

func NewVPath(path string) *VPath {
	basis := NewPath(path)
	return &VPath{basis, generateVisible(basis.General)}
}

func (p *VPath) regenerateVisible() {
	p.Visible = generateVisible(p.General)
}

func (p *VPath) Formatted() string {
	formatted := p.Visible
	cUser, err := user.Current()
	if err == nil {
		homeDir := NewVPath(cUser.HomeDir)
		formatted = strings.Replace(formatted, homeDir.Visible, "~", 1)
	}
	return formatted
}

func (p *VPath) CurrentDir() string {
	dirs := strings.Split(p.Visible, "/")
	if len(dirs) == 2 {
		return "/"
	}
	return dirs[len(dirs)-1]
}

func (p *VPath) ChangeDir(relative string) error {
	err := p.Path.ChangeDir(relative)
	if err != nil {
		return err
	}
	p.regenerateVisible()
	return nil
}

func generateVisible(general string) string {
	if runtime.GOOS != "windows" {
		return general
	}
	general = strings.Replace(general, ":", "", 1)
	return "/" + strings.TrimSuffix(general, "/")
}

func generateGeneral(visible string) string {
	if runtime.GOOS != "windows" {
		return visible
	}
	general := strings.TrimPrefix(visible, "/")
	if len(strings.Split(general, "/")) == 1 {
		return general + ":/"
	}
	return strings.Replace(general, "/", ":/", 1)
}
