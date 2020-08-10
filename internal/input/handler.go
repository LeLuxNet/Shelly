package input

import (
	"fmt"
	"github.com/LeLuxNet/Shelly/internal/sudo"
	"github.com/LeLuxNet/Shelly/pkg/engine"
	"github.com/LeLuxNet/Shelly/pkg/output"
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"io"
	"os"
	"os/user"
	"runtime"
	"strings"
)

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
	userSymbol := "$"
	if sudo.IsRoot() {
		userSymbol = "#"
	}
	prefix := output.ColorFormatting("§2§l" + username + "@" + hostname + "§9§l" +
		session.WorkingDir.Formatted() + "§r" + userSymbol + " ")
	output.Send(prefix, session.Out)
}

func delLastChars(writer io.Writer, count int) {
	for i := 0; i < count; i++ {
		output.SendRaw([]byte{8, 32, 8}, writer)
	}
}

func ReaderInput(session *sessions.Session) {
	data := make([]byte, 1024)
	_ = output.ClearScreen(session.Out)
	if !sessions.Silent {
		output.SendNl(output.Color("Shelly"+output.NEWLINE, output.ColorItalic, output.ColorFCyan), session.Out)
	}
	sendCmdPrefix(session)
	var buffer []byte
	for session.Open {
		n, err := session.In.Read(data)
		if err == io.EOF {
			session.Open = false
			break
		}
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		for i := 0; i < n; i++ {
			buffer = handleByte(data[i], buffer, session)
		}
	}
}

func handleByte(data byte, buffer []byte, session *sessions.Session) []byte {
	fmt.Print(data, ";")

	if data == 27 || data == 91 {
		return append(buffer, data)
	}

	if len(buffer) == 2 && buffer[0] == 27 && buffer[1] == 91 {
		// Arrow-Keys
		switch data {
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
	} else if data == 8 || data == 127 {
		if session.InputStringBack() {
			entry := session.GetHistoryEntry()
			end := entry[session.InputStringPos+1:]
			session.SetHistoryEntry(entry[:session.InputStringPos] + end)
			if data == 8 && !session.EchoInput {
				output.SendRaw([]byte{32, 8}, session.Out)
			} else {
				delLastChars(session.Out, 1)
			}
			output.Send(end+" ", session.Out)
			for i := 0; i <= len(end); i++ {
				output.SendRaw([]byte{8}, session.Out)
			}
		} else if data == 8 {
			output.SendRaw([]byte{32}, session.Out)
		}
	} else if data == 13 || data == 10 {
		if data != 10 {
			output.SendRaw([]byte{10}, session.Out)
		}

		if strings.HasSuffix(session.GetHistoryEntry(), "\\") {
			output.Send("> ", session.Out)
		} else {
			engine.MultiLineInput(session.GetHistoryEntry(),
				sessions.Std{In: session.In, Out: session.Out, Err: session.Err}, session, true)
			if session.Open {
				session.HistoryPos = -1
				session.InputBuffer = ""
				session.InputStringPos = 0
				sendCmdPrefix(session)
			}
		}
	} else {
		if session.EchoInput {
			output.SendRaw([]byte{data}, session.Out)
		}
		end := session.GetHistoryEntry()[session.InputStringPos:]
		session.SetHistoryEntry(session.GetHistoryEntry()[:session.InputStringPos] + string(data) + end)
		output.Send(end, session.Out)
		for i := 0; i < len(end); i++ {
			output.SendRaw([]byte{8}, session.Out)
		}
		session.InputStringPos++
	}

	return []byte{}
}
