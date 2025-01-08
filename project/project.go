package project

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func InitializeProject(location string, ignored []string) *tview.TreeView {
	root := AddDirectory(tview.NewTreeNode(location), location, ignored, true)

	project := tview.
		NewTreeView().
		SetRoot(root).
		SetCurrentNode(root).
		SetGraphics(false).
		SetTopLevel(1)

	project.
		SetTitle("  Project ").
		SetTitleAlign(tview.AlignTop).
		SetBorderPadding(1, 0, 1, 0).
		SetBorder(true)

	return project
}

func CreateModal(app *tview.Application, project *tview.TreeView) *tview.Flex {
	input := Input(
		func(textToCheck string, lastChar rune) bool {
			return true
		},
		func(event *tcell.EventKey) *tcell.EventKey {
			return event
		},
	)

	create := CreateButton(
		func(event *tcell.EventKey) *tcell.EventKey {
			return event
		},
	)

	form := tview.NewGrid().
		SetRows(0, 3).
		SetColumns(0).
		AddItem(input, 0, 0, 1, 1, 0, 0, true).
		AddItem(create, 1, 0, 1, 1, 0, 0, false)

	form.
		SetBorder(true).
		SetTitle(" Add File ").
		SetTitleAlign(tview.AlignLeft).
		SetBorderPadding(0, 0, 1, 1).
		SetBackgroundColor(tcell.ColorBlack)

	alignment := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(form, 8, 1, true).
		AddItem(nil, 0, 1, false)

	modal := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(alignment, 0, 1, true).
		AddItem(nil, 0, 1, false)

	return modal
}

func AddDirectory(root *tview.TreeNode, location string, ignored []string, isRoot bool) *tview.TreeNode {
	directory := CreateDirectory(location, ignored, isRoot)

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
