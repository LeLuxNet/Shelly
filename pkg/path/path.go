package path

import (
	"io/ioutil"
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
		target = path.Join(base,
			strings.ReplaceAll(relative, "\\", "/"))
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

func (p *Path) MkDir(name string, mayExists bool) error {
	dir, err := getRelativePathString(p.General, name, false)
	if err != nil {
		return err
	}
	err = os.Mkdir(dir, os.ModeDir)
	if mayExists && os.IsExist(err) {
		return nil
	}
	return err
}
