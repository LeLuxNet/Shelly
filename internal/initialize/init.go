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
	cUser, err := user.Current()
	if err != nil {
		fmt.Println("Unable to get cUser homedir: " + err.Error())
		os.Exit(1)
	}

	folder, err := path.NewPath(cUser.HomeDir).GetRelativePath(ShellyUserFolder, false)
	if err != nil {
		fmt.Println("Unable to get dir: " + err.Error())
		os.Exit(1)
	}

	err = folder.MkDir(true)
	if err != nil {
		fmt.Println("Unable to create dir: " + err.Error())
		os.Exit(1)
	}

	subFolder, err := folder.GetRelativePath(ScriptsFolder, false)
	if err != nil {
		fmt.Println("Unable to get dir: " + err.Error())
		os.Exit(1)
	}
	err = subFolder.MkDir(true)
	if err != nil {
		fmt.Println("Unable to create dir: " + err.Error())
		os.Exit(1)
	}
}
