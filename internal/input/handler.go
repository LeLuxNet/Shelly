package input

import (
	"fmt"
	"github.com/LeLuxNet/Shelly/pkg/engine"
	"github.com/LeLuxNet/Shelly/pkg/output"
	"github.com/LeLuxNet/Shelly/pkg/session"
	"io"
	"os"
	"os/user"
	"regexp"
	"runtime"
	"strings"
)

const CmdPrefix = "%s%s:%s$ "

func sendCmdPrefix(session *session.Session) {
	cUser, err := user.Current()
	username := ""
	if err == nil {
		username = cUser.Username
		if runtime.GOOS == "windows" {
			username = strings.Split(username, "\\")[1]
		}
	}
	hostname, _ := os.Hostname()
	prefix := fmt.Sprintf(CmdPrefix, output.Color(username, output.COLOR_F_GREEN),
		output.Color("@"+hostname, output.COLOR_F_GREEN), output.Color(session.WorkingDir.Formatted(), output.COLOR_FB_BLUE))
	output.Send(prefix, session.Out)
}

func delLastChars(writer io.Writer, count int) {
	for i := 0; i < count; i++ {
		output.SendRaw([]byte{8, 32, 8}, writer)
	}
}

func ReaderInput(session *session.Session) {
	data := make([]byte, 1024)
	sendCmdPrefix(session)
	for session.Open {
		n, err := session.In.Read(data)
		if err == io.EOF {
			err := session.Out.Close()
			if err != nil {
				fmt.Println("Error closing:", err.Error())
			}
			break
		}
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		raw := data[:n]
		text := string(raw)
		// fmt.Print(raw)
		if len(raw) == 3 && raw[0] == 27 && raw[1] == 91 {
			switch raw[2] {
			case 65:
				delLastChars(session.Out, len(session.GetHistoryEntry()))
				session.HistoryPast()
				output.Send(session.GetHistoryEntry(), session.Out)
			case 66:
				delLastChars(session.Out, len(session.GetHistoryEntry()))
				session.HistoryPresent()
				output.Send(session.GetHistoryEntry(), session.Out)
			case 67: // Right
			case 68: // Left
			}
		} else if len(raw) == 1 && (raw[0] == 8 || raw[0] == 127) {
			if len(session.GetHistoryEntry()) > 0 {
				entry := session.GetHistoryEntry()
				session.SetHistoryEntry(entry[:len(entry)-1])
				if raw[0] == 8 && !session.EchoInput {
					output.SendRaw([]byte{32, 8}, session.Out)
				} else {
					delLastChars(session.Out, 1)
				}
			} else if raw[0] == 8 {
				output.SendRaw([]byte{32}, session.Out)
			}
		} else if len(raw) == 1 && raw[0] == 9 {
			// Tab
		} else {
			if session.EchoInput {
				output.SendRaw(raw, session.Out)
			}
			session.SetHistoryEntry(session.GetHistoryEntry() + text)
		}
		if !strings.Contains(session.GetHistoryEntry(), "\n") && strings.Contains(session.GetHistoryEntry(), "\r") {
			output.SendRaw([]byte{10}, session.Out)
		}
		regex := regexp.MustCompile(`\r\n|\r\x00|\r`)
		session.SetHistoryEntry(regex.ReplaceAllString(session.GetHistoryEntry(), "\n"))
		if strings.HasSuffix(session.GetHistoryEntry(), "\\\n") {
			output.Send("> ", session.Out)
		} else if strings.HasSuffix(session.GetHistoryEntry(), "\n") {
			engine.MultiLineInput(session.GetHistoryEntry(), session)
			session.HistoryPos = -1
			session.InputBuffer = ""
			sendCmdPrefix(session)
		}
	}
}
