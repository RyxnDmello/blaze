package controls

import (
	PROJECT "blaze/project"
	"blaze/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func InitializeControls(app *tview.Application, pages *tview.Pages, project *tview.TreeView, projectModal *tview.Flex, editor *tview.TextArea) {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlE:
			reference, ok := project.GetCurrentNode().GetReference().(*PROJECT.Node)

			if !ok || reference.IsDir() {
				break
			}

			app.SetFocus(editor)

		case tcell.KeyCtrlP:
			app.SetFocus(project)

		case tcell.KeyCtrlX:
			reference, ok := project.GetCurrentNode().GetReference().(*PROJECT.Node)

			if !ok || !reference.IsDir() {
				break
			}

			if projectModal.HasFocus() {
				pages.SendToBack("MODAL").HidePage("MODAL")
				break
			}

			pages.SendToFront("MODAL").ShowPage("MODAL")

		case tcell.KeyEsc:
			if projectModal.HasFocus() {
				pages.SendToBack("MODAL").HidePage("MODAL")
				break
			}

			utils.Exit(app, "Terminated Successfully")
		}

		return event
	})
}
