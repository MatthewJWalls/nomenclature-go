package util

import (
	"os"
	"testing"
	"io/ioutil"
)

func TestGetFileLines(t *testing.T) {

	ioutil.WriteFile(
		".ignore",
		[]byte("one\ntwo\nthree\n\n\n"),
		0644,
	)

	expected  := []string{"one", "two", "three"}	
	linesRead := GetFileLines(".ignore")

	if len(linesRead) != 3 {
		t.Error("Wrong number of lines read from file")
	}

	for i := 0; i < len(expected); i += 1 {
		if linesRead[i] != expected[i] {
			t.Errorf("Invalid line. Expected %q", expected[i])
		}
	}

	os.Remove(".ignore")

}
