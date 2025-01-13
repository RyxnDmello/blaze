package controls

import (
	PROJECT "blaze/project"

	"github.com/rivo/tview"
)

func EditorControls(app *tview.Application, project *tview.TreeView, editor *tview.TextArea) {
	reference, ok := project.GetCurrentNode().GetReference().(*PROJECT.Node)

	if !ok || reference.IsDir() {
		return
	}

	if editor.HasFocus() {
		app.SetFocus(project)
		return
	}

	app.SetFocus(editor)
}
