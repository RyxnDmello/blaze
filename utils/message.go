package utils

import (
	"fmt"
	"os"

	"github.com/rivo/tview"
)

func Error(app *tview.Application, err interface{}) {
	if app != nil {
		app.Stop()
	}

	fmt.Printf("\n\n󰈸 Blaze CLI ::: %s%s%s\n\n", "\033[31m", err.(string), "\033[0m")
	os.Exit(0)
}

func Exit(app *tview.Application, message string) {
	if app != nil {
		app.Stop()
	}

	fmt.Printf("\n\n󰈸 Blaze CLI ::: %s%s%s\n\n", "\033[32m", message, "\033[0m")
	os.Exit(0)
}
