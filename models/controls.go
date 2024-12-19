package models

import (
	"blaze/directory"
	"blaze/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Controls(app *tview.Application, project *tview.TreeView, editor *tview.TextArea) {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlE:
			reference, ok := project.GetCurrentNode().GetReference().(*directory.Node)

			if !ok || reference.IsDir() {
				break
			}

			app.SetFocus(editor)

		case tcell.KeyCtrlP:
			app.SetFocus(project)

		case tcell.KeyEsc:
			utils.Exit(app, "Terminated Successfully")
		}

		return event
	})
}
