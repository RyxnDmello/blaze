package main

import (
	"fmt"
	"strings"

	"github.com/blaze/directory"
)

func main() {
	tree := directory.CreateDirectory("../../Next/Lightning")

	Expanded(tree, 0)
}

func Expanded(tree directory.Directory, spaces int) {
	tree = directory.Ignore(tree, []string{".git", ".vscode", "node_modules", ".next"})
	tree = directory.Order(tree)

	for i := 0; i < len(tree); i++ {
		fmt.Printf("%s%s %s\n", strings.Repeat(" ", spaces), tree[i].Icon(), tree[i].Name())

		if !tree[i].IsDir() {
			continue
		}

		Expanded(*tree[i].Children(), spaces+2)
	}
}
