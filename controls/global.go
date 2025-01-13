package controls

import (
	"blaze/utils"

	"github.com/rivo/tview"
)

func EscapeControls(app *tview.Application, pages *tview.Pages, createNodeModal *tview.Flex, deleteNodeModal *tview.Flex) {
	if createNodeModal.HasFocus() {
		pages.SendToBack("CREATE_NODE_MODAL").HidePage("CREATE_NODE_MODAL")
		return
	}

	if deleteNodeModal.HasFocus() {
		pages.SendToBack("DELETE_NODE_MODAL").HidePage("DELETE_NODE_MODAL")
		return
	}

	utils.Exit(app, "Terminated Successfully")
}
