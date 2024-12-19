package main

import (
	"blaze/cmd"
	"blaze/utils"

	"github.com/rivo/tview"
)

var app = tview.NewApplication()

func main() {
	defer func() {
		if err := recover(); err != nil {
			utils.Error(app, err)
		}
	}()

	layout := cmd.Blaze(app)

	if err := app.SetRoot(layout, true).Run(); err != nil {
		panic("An Unexpected Error Has Occurred")
	}
}
