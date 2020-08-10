package output

import (
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"strconv"
	"strings"
)

const (
	ColorReset         = 0
	ColorBold          = 1
	ColorItalic        = 3
	ColorUnderline     = 4
	ColorStrikethrough = 9
)

const (
	ColorFBlack = iota + 30
	ColorFRed
	ColorFGreen
	ColorFYellow
	ColorFBlue
	ColorFMagenta
	ColorFCyan
	ColorFWhite
)

const (
	ColorFBBlack = iota + 90
	ColorFBRed
	ColorFBGreen
	ColorFBYellow
	ColorFBBlue
	ColorFBMagenta
	ColorFBCyan
	ColorFBWhite
)

const (
	ColorBBlack = iota + 40
	ColorBRed
	ColorBGreen
	ColorBYellow
	ColorBBlue
	ColorBMagenta
	ColorBCyan
	ColorBWhite
)

const (
	ColorBBBlack = iota + 100
	ColorBBRed
	ColorBBGreen
	ColorBBYellow
	ColorBBBlue
	ColorBBMagenta
	ColorBBCyan
	ColorBBWhite
)

var ColorsFRainbow = [...]int{
	ColorFGreen,
	ColorFBGreen,
	ColorFBYellow,
	ColorFYellow,
	ColorFRed,
	ColorFBRed,
	ColorFBMagenta,
	ColorFMagenta,
	ColorFBlue,
	ColorFBBlue,
	ColorFCyan,
	ColorFBCyan,
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
	return GetColor(colors...) + text + GetColor(ColorReset)
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
		return ColorFBlack
	case '1':
		return ColorFBlue
	case '2':
		return ColorFGreen
	case '3':
		return ColorFCyan
	case '4':
		return ColorFRed
	case '5':
		return ColorFMagenta
	case '6':
		return ColorFYellow
	case '7':
		return ColorFWhite
	case '8':
		return ColorFBBlack
	case '9':
		return ColorFBBlue
	case 'a':
		return ColorFBGreen
	case 'b':
		return ColorFBCyan
	case 'c':
		return ColorFBRed
	case 'd':
		return ColorFBMagenta
	case 'e':
		return ColorFBYellow
	case 'f':
		return ColorFBWhite
	case 'k':
		// Obfuscated
		return 0
	case 'l':
		return ColorBold
	case 'm':
		return ColorStrikethrough
	case 'n':
		return ColorUnderline
	case 'o':
		return ColorItalic
	}
	return ColorReset
}
