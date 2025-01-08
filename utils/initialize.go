package utils

import (
	"os"
)

func Initialize() string {
	path, err := os.Getwd()

	if err != nil {
		panic("Failed To Initialize Blaze")
	}

	if !valid(path) {
		panic("Not A Git Repository")
	}

	return path
}

func valid(path string) bool {
	items, err := os.ReadDir(path)

	if err != nil {
		panic("Not A Valid Folder")
	}

	for _, item := range items {
		if item.Name() == ".git" {
			return true
		}
	}

	return false
}
