package utils

import (
	"io/fs"
	"os"
)

func GetItems(path string) []fs.DirEntry {
	items, _ := os.ReadDir(path)
	return items
}
