package main

import (
	"testing"
)

func TestPath(t *testing.T) {
	path := NewPath("C:\\Users\\test")

	general := "C:/Users/test"
	if path.General != general {
		t.Errorf("General path should be %s but is %s ", general, path.General)
	}
	visible := "/C/Users/test"
	if path.General != general {
		t.Errorf("Visible path should be %s but is %s ", visible, path.Visible)
	}

	notExistingDir := "dontExists"
	_, err := path.GetRelativePath(notExistingDir)
	if err == nil {
		t.Errorf("The dir %s/%s should not exist but got no error (GetRelativePath)", path.General, notExistingDir)
	}
	if path.ChangeDir(notExistingDir) == nil {
		t.Errorf("The dir %s/%s should not exist but got no error (ChangeDir)", path.General, notExistingDir)
	}
}
