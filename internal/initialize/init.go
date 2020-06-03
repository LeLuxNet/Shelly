package initialize

import (
	"fmt"
	"github.com/LeLuxNet/Shelly/internal/command_inpl"
	"github.com/LeLuxNet/Shelly/pkg/path"
	"os"
	"os/user"
)

const (
	RunningEnv = "SHELLY_RUNNING"

	ShellyUserFolder = ".shelly"
	ScriptsFolder    = "scripts"
)

func Init() {
	createUserDir()

	command_inpl.Register()

	_ = os.Setenv(RunningEnv, "1")
}

func createUserDir() {
	user, err := user.Current()
	if err != nil {
		fmt.Println("Unable to get user homedir: " + err.Error())
		os.Exit(1)
	}

	folder := path.NewPath(user.HomeDir)

	err = folder.MkDir(ShellyUserFolder, true)
	if err != nil {
		fmt.Println("Unable to create dir: " + err.Error())
		os.Exit(1)
	}
	err = folder.ChangeDir(ShellyUserFolder)
	if err != nil {
		fmt.Println("Unable to change dir: " + err.Error())
		os.Exit(1)
	}

	err = folder.MkDir(ScriptsFolder, true)
	if err != nil {
		fmt.Println("Unable to create dir: " + err.Error())
		os.Exit(1)
	}
	err = folder.ChangeDir(ScriptsFolder)
}
