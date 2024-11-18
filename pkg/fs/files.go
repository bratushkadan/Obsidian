package fs

import (
	"fmt"
	"os"
	"path"
)

func listRecursive(dirPath string) ([]string, error) {
	var items []string

	dirContents, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, dirent := range dirContents {
		direntPath := path.Join(dirPath, dirent.Name())
		if dirent.IsDir() {
			subdirItems, err := ListRecursive(direntPath)
			if err != nil {
				return nil, err
			}
			items = append(items, subdirItems...)
			continue
		}
		items = append(items, direntPath)
	}

	return items, nil
}

func ListRecursive(dirPath string) ([]string, error) {
	contents, err := listRecursive(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to list directory %s recursively: %w", dirPath, err)
	}
	return contents, err
}
