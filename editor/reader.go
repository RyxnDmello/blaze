package editor

import (
	"fmt"
	"os"
)

func Reader(path string, name string, isDir bool) (string, string) {
	if isDir {
		return fmt.Sprintf(" \ue5ff %s ", name), ""
	}

	text, err := os.ReadFile(path)

	if err != nil {
		panic("Not A File")
	}

	return fmt.Sprintf(" \uf016 %s ", name), string(text)
}
