package controls

import (
	PROJECT "blaze/project"

	"github.com/rivo/tview"
)

func ProjectControls(app *tview.Application, project *tview.TreeView, editor *tview.TextArea) {
	reference, ok := project.GetCurrentNode().GetReference().(*PROJECT.Node)

	if !ok || reference.IsDir() {
		return
	}

	if project.HasFocus() {
		app.SetFocus(editor)
		return
	}

	app.SetFocus(project)
}

func CreateNodeModalControls(project *tview.TreeView, pages *tview.Pages, createNodeModal *tview.Flex) {
	reference, ok := project.GetCurrentNode().GetReference().(*PROJECT.Node)

	if !ok || !reference.IsDir() {
		return
	}

	if createNodeModal.HasFocus() {
		pages.SendToBack("CREATE_NODE_MODAL").HidePage("CREATE_NODE_MODAL")
		return
	}

	pages.SendToFront("CREATE_NODE_MODAL").ShowPage("CREATE_NODE_MODAL")
}

func DeleteNodeModalControls(project *tview.TreeView, pages *tview.Pages, deleteNodeModal *tview.Flex) {
	if deleteNodeModal.HasFocus() {
		pages.SendToBack("DELETE_NODE_MODAL").HidePage("DELETE_NODE_MODAL")
		return
	}

	pages.SendToFront("DELETE_NODE_MODAL").ShowPage("DELETE_NODE_MODAL")
}
