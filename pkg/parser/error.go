package parser

const (
	NotParsableType = iota
	NotInt
)

type ParseError struct {
	Reason int
}

func (ParseError) Code() int {
	return 5
}

func (e ParseError) Error() string {
	switch e.Reason {
	case NotParsableType:
		return "This type is not parsable"
	case NotInt:
		return "Argument has to be a number"
	}
	return "Parse Error"
}
