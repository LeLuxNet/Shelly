package path

const (
	NotExists = iota + 1
	NoDir
	NoFile
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
	case NoDir:
		return "Not a directory"
	case NoFile:
		return "Is a directory"
	}
	return "Unknown Path Error"
}
