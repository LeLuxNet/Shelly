package main

import "testing"

func TestColors(t *testing.T) {
	if GetColor(COLOR_F_BLACK, COLOR_BB_WHITE) != "\u001b[30;107m" {
		t.Errorf("Color code for black foreground and bright white background is wrong")
	}
	if Color("ColorTest", COLOR_F_BLUE) != "\u001b[34mColorTest\u001b[0m" {
		t.Errorf("Text color is wrong or reset is missing")
	}
}
