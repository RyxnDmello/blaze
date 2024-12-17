package cmd

import (
	"blaze/directory"
	"blaze/models"

	"github.com/rivo/tview"
)

var (
	project *tview.TreeView
	editor  *tview.TextArea
)

var ignored = []string{".vscode", ".git", ".next", "node_modules"}

func Blaze() *tview.Grid {
	project = models.Project("../../Next/RyxnPortfolio", ignored)
	editor = models.Editor()

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

		if !ok || reference.IsDir() {
			return
		}

		models.SetTitle(editor, reference.Name())
		editor.SetText(models.Writer(reference.Path()), false)
	})

	blaze := tview.NewGrid().
		AddItem(project, 0, 0, 1, 1, 0, 0, true).
		AddItem(editor, 0, 1, 1, 1, 0, 0, false)

	return blaze
}
