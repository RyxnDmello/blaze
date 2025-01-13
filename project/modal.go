package project

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Preview(text string) *tview.TextView {
	preview := tview.
		NewTextView().
		SetText(text).
		SetTextAlign(tview.AlignCenter)

	preview.SetBorder(true)

	return preview
}

func Input(handleAccept func(textToCheck string, lastChar rune) bool, handleChange func(event *tcell.EventKey) *tcell.EventKey) *tview.InputField {
	input := tview.
		NewInputField().
		SetPlaceholder("Name...").
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetPlaceholderStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorGray)).
		SetAcceptanceFunc(handleAccept)

	input.
		SetBorder(true).
		SetBorderPadding(0, 0, 1, 1).
		SetInputCapture(handleChange)

	return input
}

func CreateButton(accept func(event *tcell.EventKey) *tcell.EventKey) *tview.Button {
	button := tview.
		NewButton("Create").
		SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack)).
		SetActivatedStyle(tcell.StyleDefault.Background(tcell.ColorBlack))

	button.
		SetBorder(true).
		SetInputCapture(accept)

	return button
}

func DeleteButton(accept func(event *tcell.EventKey) *tcell.EventKey) *tview.Button {
	button := tview.
		NewButton("Create").
		SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack)).
		SetActivatedStyle(tcell.StyleDefault.Background(tcell.ColorBlack))

	button.
		SetBorder(true).
		SetInputCapture(accept)

	return button
}
