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
			utils.Error(err)
		}
	}()

	blaze := cmd.Blaze()

	if err := app.SetRoot(blaze, true).Run(); err != nil {
		panic("An Unexpected Error Has Occurred")
	}
}
