package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bratushkadan/obsidian/pkg/fs"
	"github.com/bratushkadan/obsidian/pkg/obsidian"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	filePaths, err := fs.ListRecursive("/Users/bratushkadan/Obsidian/bratushkadan/Work/pomodoro/")
	if err != nil {
		return err
	}

	fmt.Println(strings.Join(filePaths, "\n"))
	fmt.Println()

	for _, filePath := range filePaths {
		fileContents, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		fm, err := obsidian.ParseFrontmatter(fileContents)
		if err != nil {
			return err
		}

		fm["filepath"] = filePath
		_ = json.NewEncoder(os.Stdout).Encode(fm)
	}

	return nil
}
