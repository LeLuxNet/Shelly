package path

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
