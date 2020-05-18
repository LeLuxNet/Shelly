package command_inpl

import (
	"github.com/LeLuxNet/Shelly/pkg/command"
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"image"
	"image/jpeg"
	"log"
	"os"
	"time"
)

var Time = time.Unix(0, 0)

type Anf struct{}

func (Anf) Run(args []string, std sessions.Std, session *sessions.Session) error {
	if len(args) != 2 {
		return command.WrongArgCountError{Min: 1, Max: 1}
	}
	path, err := session.WorkingDir.GetRelativePath(args[1], true)
	if err != nil {
		return err
	}
	err = path.ExpectDir(false)
	if err != nil {
		return err
	}

	// Remove exif
	err = removeExif(path.General)
	if err != nil {
		return err
	}

	// Change access and modification times
	err = path.Times(Time)
	if err != nil {
		return err
	}
	return nil
}

func removeExif(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	file, err = os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return jpeg.Encode(file, img, nil)
}
