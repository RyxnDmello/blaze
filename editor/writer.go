package editor

import "os"

func Writer(path string, text string) {
	err := os.WriteFile(path, []byte(text), os.ModeAppend)

	if err != nil {
		return
	}
}

