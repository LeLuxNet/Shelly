package parser

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	NoPrefix    = ""
	ShortPrefix = "-"
	LongPrefix  = "--"
)

type Arg struct {
	Prefix   string
	Default  interface{}
	HasParam bool
}

func (a Arg) parse(param string) (interface{}, error) {
	switch a.Default.(type) {
	case string:
		return param, nil
	case int:
		return strconv.Atoi(param)
	case time.Duration:
		return ParseDuration(param)
	}
	return nil, ParseError{}
}

type ArgParser struct {
	args []Arg
}

func NewArgParser() ArgParser {
	return ArgParser{}
}

func (a ArgParser) AddArg(arg Arg) {
	a.args = append(a.args, arg)
}

func (a ArgParser) Parse(args []string) {
	for _, raw := range args {
		if strings.HasPrefix(raw, LongPrefix) {

		} else if strings.HasPrefix(raw, ShortPrefix) {
			chars := strings.Split(strings.TrimPrefix(raw, ShortPrefix), "")
			fmt.Println(chars)
		} else {

		}
	}
}
