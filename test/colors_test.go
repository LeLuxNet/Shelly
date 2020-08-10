package test

import (
	"github.com/LeLuxNet/Shelly/pkg/output"
	"testing"
)

func TestColors(t *testing.T) {
	if output.GetColor(output.ColorFBlack, output.ColorBBWhite) != "\u001b[30;107m" {
		t.Errorf("Color code for black foreground and bright white background is wrong")
	}
	if output.Color("ColorTest", output.ColorFBlue) != "\u001b[34mColorTest\u001b[0m" {
		t.Errorf("Text color is wrong or reset is missing")
	}
}
