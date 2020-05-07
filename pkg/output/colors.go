package output

import (
	"strconv"
	"strings"
)

const (
	COLOR_RESET     = 0
	COLOR_BOLD      = 1
	COLOR_ITALIC    = 2
	COLOR_UNDERLINE = 3
)

const (
	COLOR_F_BLACK = iota + 30
	COLOR_F_RED
	COLOR_F_GREEN
	COLOR_F_YELLOW
	COLOR_F_BLUE
	COLOR_F_MAGENTA
	COLOR_F_CYAN
	COLOR_F_WHITE
)

const (
	COLOR_FB_BLACK = iota + 90
	COLOR_FB_RED
	COLOR_FB_GREEN
	COLOR_FB_YELLOW
	COLOR_FB_BLUE
	COLOR_FB_MAGENTA
	COLOR_FB_CYAN
	COLOR_FB_WHITE
)

const (
	COLOR_B_BLACK = iota + 40
	COLOR_B_RED
	COLOR_B_GREEN
	COLOR_B_YELLOW
	COLOR_B_BLUE
	COLOR_B_MAGENTA
	COLOR_B_CYAN
	COLOR_B_WHITE
)

const (
	COLOR_BB_BLACK = iota + 100
	COLOR_BB_RED
	COLOR_BB_GREEN
	COLOR_BB_YELLOW
	COLOR_BB_BLUE
	COLOR_BB_MAGENTA
	COLOR_BB_CYAN
	COLOR_BB_WHITE
)

var COLORS_F_RAINBOW = []int{
	COLOR_F_GREEN,
	COLOR_FB_GREEN,
	COLOR_FB_YELLOW,
	COLOR_F_YELLOW,
	COLOR_F_RED,
	COLOR_FB_RED,
	COLOR_FB_MAGENTA,
	COLOR_F_MAGENTA,
	COLOR_F_BLUE,
	COLOR_FB_BLUE,
	COLOR_F_CYAN,
	COLOR_FB_CYAN,
}

func GetColor(colors ...int) string {
	var sColors []string
	for i := range colors {
		sColors = append(sColors, strconv.Itoa(colors[i]))
	}
	return "\u001b[" + strings.Join(sColors, ";") + "m"
}

func Color(text string, colors ...int) string {
	return GetColor(colors...) + text + GetColor(COLOR_RESET)
}
