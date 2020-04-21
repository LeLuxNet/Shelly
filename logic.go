package main

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

const (
	CmdPrefix = "%s $ "
	Newline   = "\r\n"
)

var history []string

func sendCmdPrefix(session *Session) {
	prefix := fmt.Sprintf(CmdPrefix, session.WorkingDir.CurrentDir())
	session.Out.Write([]byte(prefix))
}

func delLastChars(writer io.Writer, count int) {
	for i := 0; i < count; i++ {
		writer.Write([]byte{8, 32, 8})
	}
}

func ReaderInput(input InOutErr) {
	session := NewSession(input)
	data := make([]byte, 1024)
	sendCmdPrefix(session)
	for session.Open {
		n, err := session.In.Read(data)
		if err == io.EOF {
			session.Out.Close()
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
				session.Out.Write([]byte(session.GetHistoryEntry()))
			case 66:
				delLastChars(session.Out, len(session.GetHistoryEntry()))
				session.HistoryPresent()
				session.Out.Write([]byte(session.GetHistoryEntry()))
			case 67: // Right
			case 68: // Left
			}
		} else if len(raw) == 1 && (raw[0] == 8 || raw[0] == 127) {
			if len(session.GetHistoryEntry()) > 0 {
				session.inputBuffer = session.inputBuffer[:len(session.inputBuffer)-1]
				if raw[0] == 8 && !session.Echo {
					session.Out.Write([]byte{32, 8})
				} else {
					delLastChars(session.Out, 1)
				}
			} else {
				session.Out.Write([]byte{32})
			}
		} else {
			if session.Echo {
				session.Out.Write(raw)
			}
			session.inputBuffer += text
		}
		if !strings.Contains(text, "\n") && strings.Contains(text, "\r") {
			session.Out.Write([]byte{10})
		}
		regex := regexp.MustCompile(`\r\n|\r\x00|\r`)
		session.inputBuffer = regex.ReplaceAllString(session.inputBuffer, "\n")
		if strings.HasSuffix(session.inputBuffer, "\\\n") {
			session.Out.Write([]byte("> "))
		} else if strings.HasSuffix(session.inputBuffer, "\n") {
			MultiLineInput(session.GetHistoryEntry(), session)
			session.HistoryPos = -1
			session.inputBuffer = ""
			sendCmdPrefix(session)
		}
		// fmt.Println("Buffer", session.inputBuffer)
	}
}

func MultiLineInput(text string, session *Session) {
	text = strings.TrimSpace(strings.ReplaceAll(text, "\\\n", ""))
	if text == "" {
		return
	}
	regex := regexp.MustCompile(`( *[;\n] *)`)
	lines := regex.Split(text, -1)
	for _, line := range lines {
		singleLineInput(line, session)
	}
}

func singleLineInput(line string, session *Session) {
	history = append(history, line)
	ands := strings.Split(line, " && ")
	for _, and := range ands {
		code := singleCommandInput(and, session)
		if code != 0 {
			break
		}
	}
}

func singleCommandInput(cmd string, session *Session) int {
	regex := regexp.MustCompile(`\s+`)
	args := regex.Split(strings.TrimSpace(cmd), -1)
	var command Cmd = alwaysFail{err: NoCmd{}}
	if args[0] == "echo" {
		command = Echo{}
	} else if args[0] == "cat" {
		command = Cat{}
	} else if args[0] == "cd" {
		command = Cd{}
	} else if args[0] == "ls" {
		command = Ls{}
	} else if args[0] == "exit" {
		command = Exit{}
	}
	err := command.Run(args, session)
	if err != nil {
		io.WriteString(session.Err, strconv.Itoa(err.Code())+": "+err.Error()+Newline)
		return err.Code()
	}
	return 0
}
