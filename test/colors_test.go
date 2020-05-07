package test

import (
	"github.com/LeLuxNet/Shelly/pkg/output"
	"testing"
)

func TestColors(t *testing.T) {
	if output.GetColor(output.COLOR_F_BLACK, output.COLOR_BB_WHITE) != "\u001b[30;107m" {
		t.Errorf("Color code for black foreground and bright white background is wrong")
	}
	if output.Color("ColorTest", output.COLOR_F_BLUE) != "\u001b[34mColorTest\u001b[0m" {
		t.Errorf("Text color is wrong or reset is missing")
	}
}
