package initialize

import (
	"github.com/LeLuxNet/Shelly/internal/command_inpl"
	"os"
)

const (
	SHELLY_RUNNING_ENV = "SHELLY_RUNNING"
)

func Init() {
	command_inpl.Register()

	_ = os.Setenv(SHELLY_RUNNING_ENV, "1")
}

func createUserDir() {

}
