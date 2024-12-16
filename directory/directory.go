package directory

import (
	"path"

	"blaze/utils"
)

type Directory []Node

func Create(location string, ignored []string, isRoot bool) Directory {
	var directory Directory

	nodes := utils.Directory(location, ignored, isRoot)

	for _, node := range nodes {
		icon := ""

		if node.IsDir() {
			icon = "󰉋"
		}

		directory = append(directory, Node{
			icon:  icon,
			name:  node.Name(),
			path:  path.Join(location, node.Name()),
			isDir: node.IsDir(),
		})
	}

	return directory
}
