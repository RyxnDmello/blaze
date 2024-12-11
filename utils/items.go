package utils

import (
	"io/fs"
	"log"
	"os"
	"slices"
)

func ReadRepository(path string, ignored []string) []fs.DirEntry {
	items, err := os.ReadDir(path)

	if err != nil {
		log.Fatal("󰈸 Blaze CLI \t ::: Invalid Path Provided")
	}

	if !valid(items) {
		log.Fatal("󰈸 Blaze CLI \t ::: Not Git Repository")
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
	if len(ignored) == 0 {
		return items
	}

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
