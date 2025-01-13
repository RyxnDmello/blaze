package cmd

import (
	CONTROLS "blaze/controls"
	EDITOR "blaze/editor"
	PROJECT "blaze/project"
	UTILS "blaze/utils"

	"github.com/rivo/tview"
)

var ignored = []string{".vscode", ".git", ".next", "node_modules"}

var (
	path string

	project *tview.TreeView
	editor  *tview.TextArea

	pages  *tview.Pages
	layout *tview.Grid

	createNodeModal *tview.Flex
	deleteNodeModal *tview.Flex
)

func Blaze(app *tview.Application) *tview.Pages {
	path = UTILS.Initialize()

	project = PROJECT.InitializeProject(path, ignored)
	editor = EDITOR.InitializeEditor()

	project.SetSelectedFunc(func(node *tview.TreeNode) {
		reference, ok := node.GetReference().(*PROJECT.Node)

		if !ok {
			return
		}

		if node.IsExpanded() {
			node.SetExpanded(false).ClearChildren()
			return
		}

		if !node.IsExpanded() {
			PROJECT.AddDirectory(node, reference.Path(), ignored, false)
			node.SetExpanded(true)
		}
	})

	project.SetChangedFunc(func(node *tview.TreeNode) {
		reference, ok := node.GetReference().(*PROJECT.Node)

		if !ok {
			return
		}

		EDITOR.Edit(editor, reference.Path(), reference.Name(), reference.IsDir())
	})

	createNodeModal = PROJECT.CreateNodeModal(app, project)
	deleteNodeModal = PROJECT.DeleteNodeModal(app, project)

	layout = tview.NewGrid().
		AddItem(project, 0, 0, 1, 1, 1, 1, true).
		AddItem(editor, 0, 1, 1, 1, 1, 1, true)

	pages = tview.NewPages().
		AddPage("CREATE_NODE_MODAL", createNodeModal, true, false).
		AddPage("DELETE_NODE_MODAL", deleteNodeModal, true, false).
		AddPage("MAIN", layout, true, true)

	CONTROLS.InitializeControls(
		app,
		pages,
		project,
		editor,
		createNodeModal,
		deleteNodeModal,
	)

	return pages
}
