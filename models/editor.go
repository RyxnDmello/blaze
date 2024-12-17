package models

import (
	"os"

	"github.com/rivo/tview"
)

func Editor() *tview.TextArea {
	editor := tview.
		NewTextArea().
		SetPlaceholder("No File Detected")

	editor.
		SetBorder(true).
		SetBorderPadding(1, 1, 1, 1)

	return editor
}

func SetTitle(editor *tview.TextArea, title string) {
	editor.SetTitle(title).SetTitleAlign(tview.AlignLeft)
}

func Writer(path string) string {
	text, err := os.ReadFile(path)

	if err != nil {
		panic("Not A File")
	}

	return string(text)
}
