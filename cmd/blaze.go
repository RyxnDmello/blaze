package cmd

import (
	"blaze/directory"
	"blaze/models"

	"github.com/rivo/tview"
)

var ignored = []string{".vscode", ".git", ".next", "node_modules"}

func Blaze(app *tview.Application) *tview.Grid {
	path := models.Initialize()
	project := models.Project(path, ignored)
	editor := models.Editor()

	project.SetSelectedFunc(func(node *tview.TreeNode) {
		reference, ok := node.GetReference().(*directory.Node)

		if !ok {
			return
		}

		children := node.GetChildren()

		if len(children) == 0 {
			models.Add(node, reference.Path(), ignored, false)
			return
		}

		node.SetExpanded(!node.IsExpanded())
	})

	project.SetChangedFunc(func(node *tview.TreeNode) {
		reference, ok := node.GetReference().(*directory.Node)

		if !ok {
			return
		}

		models.Edit(editor, reference.Path(), reference.Name(), reference.IsDir())
	})

	layout := tview.NewGrid().
		AddItem(project, 0, 0, 1, 1, 1, 1, true).
		AddItem(editor, 0, 1, 1, 1, 1, 1, false)

	models.Controls(app, project, editor)

	return layout
}
