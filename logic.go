package main

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

const (
	CmdPrefix = "%s@%s:%s$ "
	Newline   = "\r\n"
)

var history []string

func sendCmdPrefix(session *Session) {
	cUser, err := user.Current()
	username := ""
	if err == nil {
		username = cUser.Username
		if runtime.GOOS == "windows" {
			username = strings.Split(username, "\\")[1]
		}
	}
	hostname, _ := os.Hostname()
	prefix := fmt.Sprintf(CmdPrefix, username, hostname, session.WorkingDir.Formatted())
	Send(prefix, session.Out)
}

func delLastChars(writer io.Writer, count int) {
	for i := 0; i < count; i++ {
		SendRaw([]byte{8, 32, 8}, writer)
	}
}

func ReaderInput(input InOutErr) {
	session := NewSession(input)
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
				Send(session.GetHistoryEntry(), session.Out)
			case 66:
				delLastChars(session.Out, len(session.GetHistoryEntry()))
				session.HistoryPresent()
				Send(session.GetHistoryEntry(), session.Out)
			case 67: // Right
			case 68: // Left
			}
		} else if len(raw) == 1 && (raw[0] == 8 || raw[0] == 127) {
			if len(session.GetHistoryEntry()) > 0 {
				entry := session.GetHistoryEntry()
				session.setHistoryEntry(entry[:len(entry)-1])
				if raw[0] == 8 && !session.Echo {
					SendRaw([]byte{32, 8}, session.Out)
				} else {
					delLastChars(session.Out, 1)
				}
			} else if raw[0] == 8 {
				SendRaw([]byte{32}, session.Out)
			}
		} else if len(raw) == 1 && raw[0] == 9 {
			// Tab
		} else {
			if session.Echo {
				SendRaw(raw, session.Out)
			}
			session.setHistoryEntry(session.GetHistoryEntry() + text)
		}
		if !strings.Contains(session.GetHistoryEntry(), "\n") && strings.Contains(session.GetHistoryEntry(), "\r") {
			SendRaw([]byte{10}, session.Out)
		}
		regex := regexp.MustCompile(`\r\n|\r\x00|\r`)
		session.setHistoryEntry(regex.ReplaceAllString(session.GetHistoryEntry(), "\n"))
		if strings.HasSuffix(session.GetHistoryEntry(), "\\\n") {
			Send(" >", session.Out)
		} else if strings.HasSuffix(session.GetHistoryEntry(), "\n") {
			MultiLineInput(session.GetHistoryEntry(), session)
			session.HistoryPos = -1
			session.inputBuffer = ""
			sendCmdPrefix(session)
		}
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
	} else if args[0] == "exit" || args[0] == "quit" {
		command = Exit{}
	} else if args[0] == "pwd" {
		command = Pwd{}
	} else if args[0] == "sleep" {
		command = Sleep{}
	} else if args[0] == "clear" {
		command = Clear{}
	} else if args[0] == "run" {
		command = Run{}
	}
	err := command.Run(args, session)
	if err != nil {
		SendNl(strconv.Itoa(err.Code())+": "+err.Error(), session.Err)
		return err.Code()
	}
	return 0
}
