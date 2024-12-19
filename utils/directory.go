package utils

import (
	"io/fs"
	"os"
	"slices"
)

func Dir(path string, ignored []string, isRoot bool) []fs.DirEntry {
	items, err := os.ReadDir(path)

	if isRoot && err != nil {
		panic("Invalid Path Provided")
	}

	items = ignore(items, ignored)

	return order(items)
}

func ignore(items []fs.DirEntry, ignored []string) []fs.DirEntry {
	var filtered []fs.DirEntry

	for _, item := range items {
		if slices.Contains(ignored, item.Name()) {
			continue
		}

		filtered = append(filtered, item)
	}

	return filtered
}

func order(items []fs.DirEntry) []fs.DirEntry {
	var folders, files []fs.DirEntry

	for _, item := range items {
		if item.IsDir() {
			folders = append(folders, item)
			continue
		}

		files = append(files, item)
	}

	return append(folders, files...)
}
