package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bratushkadan/obsidian/pkg/fs"
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
	return nil
}
