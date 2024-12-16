package models

import (
	"fmt"

	"blaze/directory"

	"github.com/rivo/tview"
)

func Project(location string, ignored []string) (*tview.TreeView, *tview.TreeNode) {
	root := add(tview.NewTreeNode(location), location, ignored, true)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root).SetTopLevel(1).SetGraphics(false)

	var active *tview.TreeNode

	tree.SetChangedFunc(func(node *tview.TreeNode) {
		_, selectable := node.GetReference().(*directory.Node)

		if !selectable {
			return
		}

		active = node
	})

	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference, selectable := node.GetReference().(*directory.Node)

		if !selectable || reference == nil {
			return
		}

		children := node.GetChildren()

		if len(children) == 0 {
			add(node, reference.Path(), ignored, false)
			return
		}

		node.SetExpanded(!node.IsExpanded())
	})

	tree.SetBorder(true)

	return tree, active
}

func add(root *tview.TreeNode, location string, ignored []string, isRoot bool) *tview.TreeNode {
	directory := directory.Create(location, ignored, isRoot)

	for _, item := range directory {
		reference := item

		node := tview.
			NewTreeNode(fmt.Sprintf("%s %s", item.Icon(), item.Name())).
			SetReference(&reference).
			SetSelectable(true)

		root.AddChild(node)
	}

	return root
}
