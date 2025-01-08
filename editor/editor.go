package editor

import (
	"github.com/rivo/tview"
)

func InitializeEditor() *tview.TextArea {
	editor := tview.NewTextArea()
	return editor
}

func Edit(editor *tview.TextArea, path string, name string, isDir bool) {
	file, text := Reader(path, name, isDir)

	editor.SetText(text, false)

	editor.SetChangedFunc(func() {
		if !editor.HasFocus() {
			return
		}

		Writer(path, editor.GetText())
	})

	editor.
		SetTitle(file).
		SetBorder(true).
		SetBorderPadding(1, 0, 1, 0).
		SetTitleAlign(tview.AlignLeft)
}
