package main

import (
	"blaze/models"
	"blaze/utils"

	"github.com/rivo/tview"
)

var app = tview.NewApplication()

func main() {
	defer func() {
		if err := recover(); err != nil {
			utils.Error(err)
		}
	}()

	project, _ := models.Project("../../Next/RyxnPortfolio", []string{".vscode", ".git", ".next", "node_modules"})

	if err := app.SetRoot(project, true).Run(); err != nil {
		panic("An Unexpected Error Has Occurred")
	}
}
