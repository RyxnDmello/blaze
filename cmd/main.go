package main

import (
	"blaze/utils"
	"blaze/window"

	"github.com/rivo/tview"
)

var app = tview.NewApplication()

func main() {
	defer func() {
		if err := recover(); err != nil {
			utils.Error(err)
		}
	}()

	tree := window.Tree("../../Next/", []string{".vscode", ".git", ".next", "node_modules"})

	if err := app.SetRoot(tree, true).Run(); err != nil {
		panic(err)
	}
}
