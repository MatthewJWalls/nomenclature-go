package util

import (
	"strings"
	"io/ioutil"
)

// Returns a string array of all non-empty lines
// in a file.
func GetFileLines(path string) ([]string) {

	data, err := ioutil.ReadFile(path)

	if err != nil {
		panic("Failed to read file list " + path)
	}

	raw := strings.Split(string(data), "\n")
	out := []string{}

	for i := 0; i < len(raw); i++ {
		if strings.TrimSpace(raw[i]) != "" {
			out = append(out, strings.TrimSpace(raw[i]))
		}
	}

	return out

}
