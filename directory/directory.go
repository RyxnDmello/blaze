package directory

import (
	"blaze/utils"
	"path"
)

type Directory []Node

func Create(location string, ignored []string, isRoot bool) Directory {
	var directory Directory

	nodes := utils.ReadDir(location, ignored, isRoot)

	for _, node := range nodes {
		icon := ""

		if node.IsDir() {
			icon = "󰉋"
		}

		directory = append(directory, Node{
			icon:  icon,
			name:  node.Name(),
			isDir: node.IsDir(),
			ref:   path.Join(location, node.Name()),
		})
	}

	return directory
}
