package path

import (
	"github.com/LeLuxNet/Shelly/pkg/errors"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"runtime"
	"strings"
)

type Path struct {
	General string
	Visible string
}

func NewPath(path string) *Path {
	general := strings.ReplaceAll(path, "\\", "/")
	return &Path{General: general, Visible: generateVisible(general)}
}

func generateVisible(general string) string {
	if runtime.GOOS != "windows" {
		return general
	}
	return "/" + strings.Replace(general, ":", "", 1)
}

func (p *Path) ChangeDir(relative string) errors.CommandError {
	target, err := getRelativePathString(p.General, relative)
	if err != nil {
		return err
	}
	p.General = target
	p.regenerateVisible()
	return nil
}

func (p Path) GetRelativePath(relative string) (*Path, errors.CommandError) {
	target, err := getRelativePathString(p.General, relative)
	if err != nil {
		return nil, err
	}
	return NewPath(target), nil
}

func getRelativePathString(base string, relative string) (string, errors.CommandError) {
	target := path.Join(base, relative)
	_, err := os.Stat(target)
	if os.IsNotExist(err) {
		return "", PathError{NotExists}
	} else if err != nil {
		return "", errors.GeneralError{Message: err.Error()}
	}
	return target, nil
}

func (p *Path) CurrentDir() string {
	dirs := strings.Split(p.Visible, "/")
	if len(dirs) == 2 {
		return "/"
	}
	return dirs[len(dirs)-1]
}

func (p *Path) Formatted() string {
	formatted := p.Visible
	cUser, err := user.Current()
	if err == nil {
		homeDir := NewPath(cUser.HomeDir)
		formatted = strings.Replace(formatted, homeDir.Visible, "~", 1)
	}
	return formatted
}

func (p *Path) ListDir(showDotted bool) (list []os.FileInfo, error errors.CommandError) {
	raw, err := ioutil.ReadDir(p.General)
	if err != nil {
		return nil, errors.GeneralError{Message: err.Error()}
	}

	var files []os.FileInfo
	for _, file := range raw {
		if showDotted || !strings.HasPrefix(file.Name(), ".") {
			files = append(files, file)
		}
	}
	return files, nil
}

func (p *Path) regenerateVisible() {
	p.Visible = generateVisible(p.General)
}