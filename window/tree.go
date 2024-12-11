package window

import (
	"fmt"
	"path"

	"blaze/directory"

	"github.com/rivo/tview"
)

func Tree(location string, ignored []string) *tview.TreeView {
	root := add(tview.NewTreeNode(location).SetText("."), location, ignored)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root).SetTopLevel(1).SetGraphics(false)

	tree.SetBorder(true)

	return tree
}

func add(root *tview.TreeNode, location string, ignored []string) *tview.TreeNode {
	directory := directory.Create(location, ignored)

	for _, item := range directory {
		node := tview.
			NewTreeNode(fmt.Sprintf("%s %s", item.Icon(), item.Name())).
			SetReference(path.Join(location, item.Name())).
			SetExpanded(item.IsDir()).
			SetSelectable(true)

		root.AddChild(node)
	}

	return root
}
