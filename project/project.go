package project

import (
	"fmt"
	"os"
	"path"
	"strings"

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

func CreateNodeModal(app *tview.Application, project *tview.TreeView) *tview.Flex {
	var input *tview.InputField
	var create *tview.Button

	input = Input(
		func(textToCheck string, lastChar rune) bool {
			if parts := strings.Split(textToCheck, "."); len(parts) > 3 {
				input.SetFieldTextColor(tcell.ColorRed)
				return true
			}

			input.SetFieldTextColor(tcell.ColorWhite)

			return true
		},
		func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTab {
				app.SetFocus(create)
			}

			return event
		},
	)

	create = CreateButton(
		func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyEnter {
				reference, ok := project.GetCurrentNode().GetReference().(*Node)

				if !ok {
					input.SetFieldTextColor(tcell.ColorRed)
					return event
				}

				path := path.Join(reference.path, input.GetText())

				_, err := os.Create(path)

				if err != nil {
					input.SetFieldTextColor(tcell.ColorRed)
					return event
				}

				input.SetText("")
			}

			if event.Key() == tcell.KeyTab {
				app.SetFocus(input)
			}

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

func DeleteNodeModal(app *tview.Application, project *tview.TreeView) *tview.Flex {
	var preview *tview.TextView
	var delete *tview.Button

	preview = Preview("Do you want to proceed?")

	delete = DeleteButton(
		func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyEnter {
				reference, ok := project.GetCurrentNode().GetReference().(*Node)

				if !ok {
					panic("Crashed Unexpectedly")
				}

				err := os.Remove(reference.path)

				if err != nil {
					panic("Crashed Unexpectedly")
				}
			}

			return event
		},
	)

	form := tview.NewGrid().
		SetRows(0, 3).
		SetColumns(0).
		AddItem(preview, 0, 0, 1, 1, 0, 0, false).
		AddItem(delete, 1, 0, 1, 1, 0, 0, true)

	form.
		SetBorder(true).
		SetTitle(" Delete File ").
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
			SetSelectable(true).
			SetExpanded(false)

		root.AddChild(node)
	}

	return root
}
