package session

import (
	"github.com/LeLuxNet/Shelly/pkg/path"
	"io"
	"os"
)

var history []string

func AddHistory(line string) {
	history = append(history, line)
}

type Session struct {
	In          io.Reader
	Out         io.WriteCloser
	Err         io.Writer
	EchoInput   bool
	WorkingDir  *path.Path
	Open        bool
	HistoryPos  int
	InputBuffer string
}

func NewSession(In io.Reader, Out io.WriteCloser, Err io.Writer, EchoInput bool) *Session {
	dir, _ := os.Getwd()
	return &Session{In: In, Out: Out, Err: Err, EchoInput: EchoInput, WorkingDir: path.NewPath(dir), Open: true, HistoryPos: -1}
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
