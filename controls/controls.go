package controls

import (
	PROJECT "blaze/project"
	"blaze/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func InitializeControls(
	app *tview.Application,
	pages *tview.Pages,
	project *tview.TreeView,
	editor *tview.TextArea,
	createNodeModal *tview.Flex,
	deleteNodeModal *tview.Flex,
) {
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

			if createNodeModal.HasFocus() {
				pages.SendToBack("CREATE_NODE_MODAL").HidePage("CREATE_NODE_MODAL")
				break
			}

			pages.SendToFront("CREATE_NODE_MODAL").ShowPage("CREATE_NODE_MODAL")

		case tcell.KeyCtrlD:
			if deleteNodeModal.HasFocus() {
				pages.SendToBack("DELETE_NODE_MODAL").HidePage("DELETE_NODE_MODAL")
				break
			}

			pages.SendToFront("DELETE_NODE_MODAL").ShowPage("DELETE_NODE_MODAL")

		case tcell.KeyEsc:
			if createNodeModal.HasFocus() {
				pages.SendToBack("CREATE_NODE_MODAL").HidePage("CREATE_NODE_MODAL")
				break
			}

			utils.Exit(app, "Terminated Successfully")
		}

		return event
	})
}
