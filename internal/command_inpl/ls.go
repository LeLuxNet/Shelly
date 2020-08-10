package command_inpl

import (
	"github.com/LeLuxNet/Shelly/pkg/output"
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"os"
	"strings"
)

type Ls struct{}

var (
	ImageEndings = []string{".jpg", ".jpeg", ".mjpg", ".mjpeg", ".gif", ".bmp", ".pbm", ".pgm", ".ppm", ".tga",
		".xbm", ".xpm", ".tif", ".tiff", ".png", ".svg", ".svgz", ".mng", ".pcx", ".mov", ".mpg", ".mpeg", ".m2v",
		".mkv", ".webm", ".ogm", ".mp4", ".m4v", ".mp4v", ".vob", ".qt", ".nuv", ".wmv", ".asf", ".rm", ".rmvb",
		".flc", ".avi", ".fli", ".flv", ".gl", ".dl", ".xcf", ".xwd", ".yuv", ".cgm", ".emf", ".ogv", ".ogx"}
	ArchiveEndings = []string{".tar", ".tgz", ".arc", ".arj", ".taz", ".lha", ".lz4", ".lzh", ".lzma", ".tlz",
		".txz", ".tzo", ".t7z", ".zip", ".z", ".dz", ".gz", ".lrz", ".lz", ".lzo", ".xz", ".zst", ".tzst", ".bz2",
		".bz", ".tbz", ".tbz2", ".tz", ".deb", ".rpm", ".jar", ".war", ".ear", ".sar", ".rar", ".alz", ".ace",
		".zoo", ".cpio", ".7z", ".rz", ".cab", ".wim", ".swm", ".dwm", ".esd"}
	AudioEndings = []string{".aac", ".au", ".flac", ".m4a", ".mid", ".midi", ".mka", ".mp3", ".mpc", ".ogg",
		".ra", ".wav", ".oga", ".opus", ".spx", ".xspf"}
	ExecutableEndings = []string{".sy", ".exe", ".bat", ".cmd", ".msi"}
)

func (Ls) Run(args []string, std sessions.Std, session *sessions.Session) error {
	files, err := session.WorkingDir.ListDir(false)
	if err != nil {
		return err
	}

	for _, file := range files {
		path, err := session.WorkingDir.GetRelativePath(file.Name(), true)
		if err != nil {
			return err
		}
		var result string
		if path.ExpectDir(true) == nil {
			result = output.GetColor(output.ColorBold, output.ColorFBlue)
		} else if file.Mode() == os.ModeSymlink {
			result = output.GetColor(output.ColorBold, output.ColorFCyan)
		} else if file.Mode() == os.ModeDevice {
			result = output.GetColor(output.ColorFBYellow, output.ColorBBlack)
		} else if hasEnding(file.Name(), ImageEndings) {
			result = output.GetColor(output.ColorBold, output.ColorFMagenta)
		} else if hasEnding(file.Name(), ArchiveEndings) {
			result = output.GetColor(output.ColorBold, output.ColorFRed)
		} else if hasEnding(file.Name(), AudioEndings) {
			result = output.GetColor(output.ColorFCyan)
		} else if file.Mode()&0111 != 0 {
			result = output.GetColor(output.ColorBold, output.ColorFGreen)
		} else if hasEnding(file.Name(), ExecutableEndings) {
			result = output.GetColor(output.ColorFGreen)
		} else {
			result = output.GetColor(output.ColorReset)
		}
		output.SendNl(result+file.Name(), std.Out)
	}
	return nil
}

func hasEnding(name string, endings []string) bool {
	for _, ending := range endings {
		if strings.HasSuffix(name, ending) {
			return true
		}
	}
	return false
}
