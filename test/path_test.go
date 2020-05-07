package test

import (
	"github.com/LeLuxNet/Shelly/pkg/path"
	"testing"
)

func TestPath(t *testing.T) {
	testPath := path.NewPath("C:\\Users\\test")

	general := "C:/Users/test"
	if testPath.General != general {
		t.Errorf("General testPath should be %s but is %s ", general, testPath.General)
	}
	visible := "/C/Users/test"
	if testPath.General != general {
		t.Errorf("Visible testPath should be %s but is %s ", visible, testPath.Visible)
	}

	notExistingDir := "dontExists"
	_, err := testPath.GetRelativePath(notExistingDir)
	if err == nil {
		t.Errorf("The dir %s/%s should not exist but got no error (GetRelativePath)", testPath.General, notExistingDir)
	}
	if testPath.ChangeDir(notExistingDir) == nil {
		t.Errorf("The dir %s/%s should not exist but got no error (ChangeDir)", testPath.General, notExistingDir)
	}
}
