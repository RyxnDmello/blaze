package models

import (
	"fmt"

	"blaze/directory"

	"github.com/rivo/tview"
)

func Project(location string, ignored []string) *tview.TreeView {
	root := Add(tview.NewTreeNode(location), location, ignored, true)

	tree := tview.
		NewTreeView().
		SetRoot(root).
		SetCurrentNode(root).
		SetGraphics(false).
		SetTopLevel(1)

	tree.
		SetTitle("  Project ").
		SetTitleAlign(tview.AlignTop).
		SetBorderPadding(1, 0, 1, 0).
		SetBorder(true)

	return tree
}

func Add(root *tview.TreeNode, location string, ignored []string, isRoot bool) *tview.TreeNode {
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
