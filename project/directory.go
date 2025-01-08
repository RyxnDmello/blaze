package project

import (
	"path"

	"blaze/utils"
)

type Directory []Node

func CreateDirectory(location string, ignored []string, isRoot bool) Directory {
	var directory Directory

	nodes := utils.Dir(location, ignored, isRoot)

	for _, node := range nodes {
		icon := "\uf016"

		if node.IsDir() {
			icon = "\ue5ff"
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
