// +build !windows

package sudo

import "os"

func IsRoot() bool {
	return os.Geteuid() == 0
}
