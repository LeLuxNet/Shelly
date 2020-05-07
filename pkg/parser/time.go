package parser

import (
	"github.com/LeLuxNet/Shelly/pkg/errors"
	"regexp"
	"strconv"
	"time"
)

func ParseTime(input string) (time.Duration, errors.CommandError) {
	regex := regexp.MustCompile(`([0-9]+)([a-z]*)`)
	parts := regex.FindStringSubmatch(input)
	var mult time.Duration
	switch parts[2] {
	case "", "s":
		mult = time.Second
	case "m":
		mult = time.Minute
	case "h":
		mult = time.Hour
	case "d":
		mult = time.Hour * 24
	}
	if &mult == nil {
		return time.Duration(0), ParseError{}
	}
	number, err := strconv.Atoi(parts[1])
	if err != nil {
		return time.Duration(0), ParseError{Reason: NotInt}
	}
	return mult * time.Duration(number), nil
}
