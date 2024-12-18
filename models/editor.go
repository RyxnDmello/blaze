package models

import (
	"blaze/editor"

	"github.com/rivo/tview"
)

func Editor() *tview.TextArea {
	code := tview.NewTextArea()
	return code
}

func Edit(code *tview.TextArea, path string, name string, isDir bool) {
	file, text := editor.Reader(path, name, isDir)

	code.SetText(text, false)

	code.SetChangedFunc(func() {
		if !code.HasFocus() {
			return
		}

		editor.Writer(path, code.GetText())
	})

	code.SetTitle(file).SetTitleAlign(tview.AlignLeft).SetBorder(true).SetBorderPadding(1, 0, 1, 0)
}
