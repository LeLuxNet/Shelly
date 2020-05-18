package path

import (
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

func generateGeneral(visible string) string {
	if runtime.GOOS != "windows" {
		return visible
	}
	return strings.Replace(strings.TrimPrefix(visible, "/"), "/", ":/", 1)
}

func (p *Path) ChangeDir(relative string) error {
	target, err := getRelativePathString(p.General, relative, true)
	if err != nil {
		return err
	}
	old := p.General
	p.General = target
	err = p.ExpectDir(true)
	if err != nil {
		p.General = old
		return err
	}
	p.regenerateVisible()
	return nil
}

func (p Path) GetRelativePath(relative string, exists bool) (*Path, error) {
	target, err := getRelativePathString(p.General, relative, exists)
	if err != nil {
		return nil, err
	}
	return NewPath(target), nil
}

func getRelativePathString(base string, relative string, exists bool) (string, error) {
	var target string
	if strings.HasPrefix(relative, "/") {
		target = generateGeneral(relative)
	} else {
		target = path.Join(base, relative)
	}
	if exists {
		_, err := os.Stat(target)
		if os.IsNotExist(err) {
			return "", PathError{NotExists}
		} else if err != nil {
			return "", err
		}
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

func (p *Path) ListDir(showDotted bool) (list []os.FileInfo, error error) {
	raw, err := ioutil.ReadDir(p.General)
	if err != nil {
		return nil, err
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

func (p *Path) ExpectDir(dir bool) error {
	file, err := os.Stat(p.General)
	if err != nil {
		return err
	}
	if file.Mode().IsRegular() && dir {
		return PathError{Id: NoDir}
	} else if file.Mode().IsDir() && !dir {
		return PathError{Id: NoFile}
	}
	return nil
}
