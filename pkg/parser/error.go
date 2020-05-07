package parser

const (
	NotInt = iota
)

type ParseError struct {
	Reason int
}

func (ParseError) Code() int {
	return 5
}

func (e ParseError) Error() string {
	switch e.Reason {
	case NotInt:
		return "Argument has to be an int"
	}
	return "Parse Error"
}
