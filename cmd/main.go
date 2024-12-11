package main

import (
	"blaze/window"
	"log"

	"github.com/rivo/tview"
)

var app = tview.NewApplication()

func main() {
	tree := window.Tree("../../Next/Lightning", []string{".vscode", ".git", "node_modules"})

	if err := app.SetRoot(tree, true).Run(); err != nil {
		log.Fatal("🔥 Blaze: Failed")
	}
}
