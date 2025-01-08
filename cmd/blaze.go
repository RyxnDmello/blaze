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

	projectModal *tview.Flex
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

		children := node.GetChildren()

		if len(children) == 0 {
			PROJECT.AddDirectory(node, reference.Path(), ignored, false)
			return
		}

		node.SetExpanded(!node.IsExpanded())
	})

	project.SetChangedFunc(func(node *tview.TreeNode) {
		reference, ok := node.GetReference().(*PROJECT.Node)

		if !ok {
			return
		}

		EDITOR.Edit(editor, reference.Path(), reference.Name(), reference.IsDir())
	})

	projectModal = PROJECT.CreateModal(app, project)

	layout = tview.NewGrid().
		AddItem(project, 0, 0, 1, 1, 1, 1, true).
		AddItem(editor, 0, 1, 1, 1, 1, 1, true)

	pages = tview.NewPages().
		AddPage("MODAL", projectModal, true, false).
		AddPage("MAIN", layout, true, true)

	CONTROLS.InitializeControls(app, pages, project, projectModal, editor)

	return pages
}
