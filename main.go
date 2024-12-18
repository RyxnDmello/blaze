package main

import (
	"blaze/cmd"
	"blaze/utils"

	"github.com/rivo/tview"
)

var app = tview.NewApplication()

func main() {
	defer func() {
		if err := recover(); err != nil {
			utils.Error(err)
		}
	}()

	project, editor := cmd.Blaze()

	blaze := tview.NewGrid().
		AddItem(project, 0, 0, 1, 1, 1, 1, true).
		AddItem(editor, 0, 1, 1, 1, 1, 1, false)

	if err := app.SetRoot(blaze, true).Run(); err != nil {
		panic("An Unexpected Error Has Occurred")
	}
}
