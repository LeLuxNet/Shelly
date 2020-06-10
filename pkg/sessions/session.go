package sessions

import (
	"github.com/LeLuxNet/Shelly/pkg/path"
	"io"
	"os"
)

var history []string
var NoColors bool
var Silent bool

func AddHistory(line string) {
	history = append(history, line)
}

type Session struct {
	In             io.Reader
	Out            io.WriteCloser
	Err            io.Writer
	EchoInput      bool
	WorkingDir     *path.VPath
	Open           bool
	HistoryPos     int
	InputStringPos int
	InputBuffer    string
}

func NewSession(In io.Reader, Out io.WriteCloser, Err io.Writer, EchoInput bool) *Session {
	dir, _ := os.Getwd()
	if err == nil {
		err = out
	}
	return &Session{In: in, Out: out, Err: err, EchoInput: echoInput, WorkingDir: path.NewVPath(dir), Open: true, HistoryPos: -1, InputStringPos: 0}
}

func (s *Session) Close() error {
	s.Open = false
	return s.Out.Close()
}

func (s *Session) HistoryPast() {
	if s.HistoryPos == 0 || len(history) == 0 {
		return
	}
	if s.HistoryPos == -1 {
		s.HistoryPos = len(history) - 1
	} else {
		s.HistoryPos--
	}
	s.InputStringPos = len(s.GetHistoryEntry())
}

func (s *Session) HistoryPresent() {
	if s.HistoryPos == -1 {
		return
	}
	if s.HistoryPos < len(history)-1 {
		s.HistoryPos++
	} else {
		s.HistoryPos = -1
	}
	s.InputStringPos = len(s.GetHistoryEntry())
}

func (s *Session) GetHistoryEntry() string {
	if s.HistoryPos == -1 {
		return s.InputBuffer
	} else {
		return history[s.HistoryPos]
	}
}

func (s *Session) SetHistoryEntry(text string) {
	if s.HistoryPos == -1 {
		s.InputBuffer = text
	} else {
		history[s.HistoryPos] = text
	}
}

func (s *Session) InputStringBack() bool {
	if s.InputStringPos > 0 {
		s.InputStringPos--
		return true
	}
	return false
}

func (s *Session) InputStringForward() bool {
	if len(s.GetHistoryEntry()) > s.InputStringPos {
		s.InputStringPos++
		return true
	}
	return false
}
