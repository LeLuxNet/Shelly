package input

import (
	"fmt"
	"github.com/LeLuxNet/Shelly/pkg/engine"
	"github.com/LeLuxNet/Shelly/pkg/output"
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"io"
	"os"
	"os/user"
	"regexp"
	"runtime"
	"strings"
)

const CmdPrefix = "%s%s:%s$ "

func sendCmdPrefix(session *sessions.Session) {
	cUser, err := user.Current()
	username := ""
	if err == nil {
		username = cUser.Username
		if runtime.GOOS == "windows" {
			username = strings.Split(username, "\\")[1]
		}
	}
	hostname, _ := os.Hostname()
	prefix := fmt.Sprintf(CmdPrefix,
		output.Color(username, output.COLOR_BOLD, output.COLOR_F_GREEN),
		output.Color("@"+hostname, output.COLOR_BOLD, output.COLOR_F_GREEN),
		output.Color(session.WorkingDir.Formatted(), output.COLOR_BOLD, output.COLOR_FB_BLUE))
	output.Send(prefix, session.Out)
}

func delLastChars(writer io.Writer, count int) {
	for i := 0; i < count; i++ {
		output.SendRaw([]byte{8, 32, 8}, writer)
	}
}

func ReaderInput(session *sessions.Session) {
	data := make([]byte, 1024)
	output.ClearScreen(session.Out)
	if !sessions.Silent {
		output.SendNl(output.Color("Shelly"+output.NEWLINE, output.COLOR_ITALIC, output.COLOR_F_CYAN), session.Out)
	}
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
			case 67:
				if session.InputStringForward() {
					output.SendRaw([]byte{27, 91, 67}, session.Out)
				}
			case 68:
				if session.InputStringBack() {
					output.SendRaw([]byte{27, 91, 68}, session.Out)
				}
			}
		} else if len(raw) == 1 && (raw[0] == 8 || raw[0] == 127) {
			if session.InputStringBack() {
				entry := session.GetHistoryEntry()
				end := entry[session.InputStringPos+1:]
				session.SetHistoryEntry(entry[:session.InputStringPos] + end)
				if raw[0] == 8 && !session.EchoInput {
					output.SendRaw([]byte{32, 8}, session.Out)
				} else {
					delLastChars(session.Out, 1)
				}
				output.Send(end+" ", session.Out)
				for i := 0; i <= len(end); i++ {
					output.SendRaw([]byte{8}, session.Out)
				}
			} else if raw[0] == 8 {
				output.SendRaw([]byte{32}, session.Out)
			}
		} else if len(raw) == 1 && raw[0] == 9 {
			// Tab
		} else if len(raw) == 1 && raw[0] == 3 {
			// Ctrl+C
		} else {
			if session.EchoInput {
				output.SendRaw(raw, session.Out)
			}
			end := session.GetHistoryEntry()[session.InputStringPos:]
			session.SetHistoryEntry(session.GetHistoryEntry()[:session.InputStringPos] + text + end)
			output.Send(end, session.Out)
			for i := 0; i < len(end); i++ {
				output.SendRaw([]byte{8}, session.Out)
			}
			session.InputStringPos++
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
			session.InputStringPos = 0
			sendCmdPrefix(session)
		}
	}
}
