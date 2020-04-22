package main

import "strconv"

type CmdCrashError interface {
	error
	Code() int
}

type GeneralError struct {
	Message string
}

func (GeneralError) Code() int {
	return 1
}

func (e GeneralError) Error() string {
	return e.Message
}

type NoCmd struct {
	Message string
}

func (NoCmd) Code() int {
	return 2
}

func (NoCmd) Error() string {
	return "No such cmd"
}

type WrongArgCountError struct {
	min int
	max int
}

func (WrongArgCountError) Code() int {
	return 3
}

func (e WrongArgCountError) Error() string {
	if e.min != 0 && e.max != 0 {
		if e.min == e.max {
			return "You need " + strconv.Itoa(e.min) + " args"
		}
		return "You need between " + strconv.Itoa(e.min) + " and " + strconv.Itoa(e.max) + " args"
	} else if e.min != 0 {
		return "You need min " + strconv.Itoa(e.min) + " args"
	} else if e.max != 0 {
		return "You need max " + strconv.Itoa(e.max) + " args"
	}
	return "Wrong arg count"
}

const (
	NotExists = iota + 1
)

type PathError struct {
	Id int
}

func (PathError) Code() int {
	return 4
}

func (e PathError) Error() string {
	switch e.Id {
	case NotExists:
		return "File or dir does not exist"
	}
	return "Unknown Path Error"
}

const (
	NotInt = iota + 1
)

type ParseError struct {
	Id int
}

func (ParseError) Code() int {
	return 5
}

func (e ParseError) Error() string {
	switch e.Id {
	case NotInt:
		return "Argument has to be an int"
	}
	return "Parse Error"
}
