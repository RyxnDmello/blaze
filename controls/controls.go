package controls

import (
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
			EditorControls(app, project, editor)

		case tcell.KeyCtrlP:
			ProjectControls(app, project, editor)

		case tcell.KeyCtrlX:
			CreateNodeModalControls(project, pages, createNodeModal)

		case tcell.KeyCtrlD:
			DeleteNodeModalControls(project, pages, deleteNodeModal)

		case tcell.KeyEsc:
			EscapeControls(app, pages, createNodeModal, deleteNodeModal)
		}

		return event
	})
}
