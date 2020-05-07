package command

import "strconv"

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
	Min int
	Max int
}

func (WrongArgCountError) Code() int {
	return 3
}

func (e WrongArgCountError) Error() string {
	if e.Min != 0 && e.Max != 0 {
		if e.Min == e.Max {
			return "You need " + strconv.Itoa(e.Min) + " args"
		}
		return "You need between " + strconv.Itoa(e.Min) + " and " + strconv.Itoa(e.Max) + " args"
	} else if e.Min != 0 {
		return "You need min " + strconv.Itoa(e.Min) + " args"
	} else if e.Max != 0 {
		return "You need max " + strconv.Itoa(e.Max) + " args"
	}
	return "Wrong arg count"
}
