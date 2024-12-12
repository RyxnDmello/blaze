package window

import (
	"fmt"

	"blaze/directory"

	"github.com/rivo/tview"
)

func Tree(location string, ignored []string) *tview.TreeView {
	root := add(tview.NewTreeNode(location).SetText("."), location, ignored, true)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root).SetTopLevel(1).SetGraphics(false)

	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()

		if reference == nil {
			return
		}

		children := node.GetChildren()

		if len(children) == 0 {
			path := reference.(string)
			add(node, path, ignored, false)
			return
		}

		node.SetExpanded(!node.IsExpanded())
	})

	tree.SetBorder(true)

	return tree
}

func add(root *tview.TreeNode, location string, ignored []string, isRoot bool) *tview.TreeNode {
	directory := directory.Create(location, ignored, isRoot)

	for _, item := range directory {
		node := tview.
			NewTreeNode(fmt.Sprintf("%s %s", item.Icon(), item.Name())).
			SetReference(item.Reference()).
			SetSelectable(true)

		root.AddChild(node)
	}

	return root
}
