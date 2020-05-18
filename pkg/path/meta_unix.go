// +build !windows

package path

import (
	"os"
	"time"
)

func (p Path) Times(time time.Time) error {
	return os.Chtimes(p.General, time, time)
}
