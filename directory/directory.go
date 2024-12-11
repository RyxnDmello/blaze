package directory

import (
	"blaze/utils"
	"path"
)

type Directory []Node

func Create(location string, ignored []string) Directory {
	var directory Directory

	nodes := utils.ReadRepository(location, ignored)

	for _, node := range nodes {
		icon := ""

		if node.IsDir() {
			icon = "󰉋"
		}

		ref := path.Join(location, node.Name())

		directory = append(directory, Node{
			ref:   ref,
			icon:  icon,
			name:  node.Name(),
			isDir: node.IsDir(),
		})
	}

	return directory
}
