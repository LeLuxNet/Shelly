package output

import (
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"strconv"
	"strings"
)

const (
	COLOR_RESET         = 0
	COLOR_BOLD          = 1
	COLOR_ITALIC        = 3
	COLOR_UNDERLINE     = 4
	COLOR_STRIKETHROUGH = 9
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
	if sessions.NoColors {
		return ""
	}
	var sColors []string
	for i := range colors {
		sColors = append(sColors, strconv.Itoa(colors[i]))
	}
	return "\u001b[" + strings.Join(sColors, ";") + "m"
}

func ColorFormatting(text string) string {
	mode := 0
	result := text
	for i, char := range text {
		if char == '§' || char == '&' {
			if mode == 0 {
				mode = 1
			}
		} else if mode != 0 {
			code := formattingCode(char)
			if text[i-1] == '&' {
				code = toBackground(code)
			}
			color := ""
			if mode == 1 {
				color += "\u001b["
			}
			color += strconv.Itoa(code)
			if len(text) > i && (text[i+2] == '§' || text[i+2] == '&') {
				mode = 2
				color += ";"
			} else {
				mode = 0
				color += "m"
			}
			result = strings.Replace(result, string(text[i-1])+string(char),
				color, 1)
		}
	}
	return result
}

func Color(text string, colors ...int) string {
	return GetColor(colors...) + text + GetColor(COLOR_RESET)
}

func toBackground(color int) int {
	if (color >= 30 && color <= 39) ||
		(color >= 90 && color <= 97) {
		return color + 60
	}
	return color
}

func formattingCode(rune rune) int {
	switch rune {
	case '0':
		return COLOR_F_BLACK
	case '1':
		return COLOR_F_BLUE
	case '2':
		return COLOR_F_GREEN
	case '3':
		return COLOR_F_CYAN
	case '4':
		return COLOR_F_RED
	case '5':
		return COLOR_F_MAGENTA
	case '6':
		return COLOR_F_YELLOW
	case '7':
		return COLOR_F_WHITE
	case '8':
		return COLOR_FB_BLACK
	case '9':
		return COLOR_FB_BLUE
	case 'a':
		return COLOR_FB_GREEN
	case 'b':
		return COLOR_FB_CYAN
	case 'c':
		return COLOR_FB_RED
	case 'd':
		return COLOR_FB_MAGENTA
	case 'e':
		return COLOR_FB_YELLOW
	case 'f':
		return COLOR_FB_WHITE
	case 'k':
		// Obfuscated
		return 0
	case 'l':
		return COLOR_BOLD
	case 'm':
		return COLOR_STRIKETHROUGH
	case 'n':
		return COLOR_UNDERLINE
	case 'o':
		return COLOR_ITALIC
	}
	return COLOR_RESET
}
