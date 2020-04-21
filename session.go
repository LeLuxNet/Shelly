package main

import "os"

type Session struct {
	InOutErr
	WorkingDir  *Path
	Open        bool
	HistoryPos  int
	inputBuffer string
}

func NewSession(inOut InOutErr) *Session {
	dir, _ := os.Getwd()
	return &Session{InOutErr: inOut, WorkingDir: NewPath(dir), Open: true, HistoryPos: -1}
}

func (s *Session) Close() {
	s.Open = false
	s.Out.Close()
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
		return s.inputBuffer
	} else {
		return history[s.HistoryPos]
	}
}
