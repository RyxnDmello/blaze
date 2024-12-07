package directory

import (
	"path"
	"slices"

	"github.com/blaze/utils"
)

type Directory []Node

func GetDirectory(location string) Directory {
	directory := CreateDirectory(location)
	directory = SetPointers(directory)
	return directory
}

func SetPointers(directory Directory) Directory {
	for i := 0; i < len(directory); i++ {
		if i > 0 {
			directory[i].prev = &directory[i-1]
		}

		if i < len(directory)-1 {
			directory[i].next = &directory[i+1]
		}
	}

	return directory
}

func CreateDirectory(location string) Directory {
	parent := make(Directory, 0)

	items := utils.GetItems(location)

	for i := 0; i < len(items); i++ {
		icon := ""

		if items[i].IsDir() {
			icon = "󰉋"
		}

		parent = append(parent, Node{
			name:     items[i].Name(),
			icon:     icon,
			isDir:    items[i].IsDir(),
			isOpen:   false,
			children: nil,
			prev:     nil,
			next:     nil,
		})

		if items[i].IsDir() {
			children := GetDirectory(path.Join(location, parent[i].name))
			parent[i].children = &children
		}
	}

	return parent
}

func Ignore(directory Directory, ignored []string) Directory {
	filtered := make(Directory, 0)

	for _, item := range directory {
		if slices.Contains(ignored, item.name) {
			continue
		}

		filtered = append(filtered, item)
	}

	return filtered
}

func Order(directory Directory) Directory {
	folders := make(Directory, 0)
	files := make(Directory, 0)

	for _, item := range directory {
		if item.isDir {
			folders = append(folders, item)
			continue
		}

		files = append(files, item)
	}

	return append(folders, files...)
}
