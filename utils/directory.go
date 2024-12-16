package utils

import (
	"io/fs"
	"os"
	"slices"
)

func Directory(path string, ignored []string, isRoot bool) []fs.DirEntry {
	items, err := os.ReadDir(path)

	if isRoot && err != nil {
		panic("Invalid Path Provided")
	}

	if isRoot && !valid(items) {
		panic("Not Git Repository")
	}

	items = ignore(items, ignored)

	return order(items)
}

func valid(items []fs.DirEntry) bool {
	for _, item := range items {
		if item.Name() == ".git" {
			return true
		}
	}

	return false
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
