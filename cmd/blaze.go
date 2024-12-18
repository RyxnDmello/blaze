package cmd

import (
	"blaze/directory"
	"blaze/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	project *tview.TreeView
	editor  *tview.TextArea
)

var ignored = []string{".vscode", ".git", ".next", "node_modules"}

func Blaze() (*tview.TreeView, *tview.TextArea) {
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

		if !ok {
			return
		}

		node.SetTextStyle(tcell.StyleDefault.Bold(false))

		models.Edit(editor, reference.Path(), reference.Name(), reference.IsDir())
	})

	return project, editor
}
