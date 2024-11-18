package obsidian

import (
	"errors"
	"regexp"

	"gopkg.in/yaml.v3"
)

var (
	metadataSeparatorRe = regexp.MustCompile(`(?m)^---$`)
)

var (
	ErrNoFrontmatter = errors.New("frontmatter not found in the markdown page")
)

func getRawMetadataContents(fileContents string) (string, error) {
	metadataSeparatorRe.FindStringIndex(fileContents)
	return "", nil
}

func ParseFrontmatter(fileContents []byte) (map[string]any, error) {
	firstLoc := metadataSeparatorRe.FindIndex(fileContents)
	if len(firstLoc) < 2 {
		return nil, ErrNoFrontmatter
	}
	secondLoc := metadataSeparatorRe.FindIndex(fileContents[firstLoc[1]+1:])
	if len(secondLoc) < 2 {
		return nil, ErrNoFrontmatter
	}

	stringFrontmatter := fileContents[firstLoc[1]+1 : secondLoc[1]]
	var frontmatter map[string]any
	err := yaml.Unmarshal(stringFrontmatter, &frontmatter)
	if err != nil {
		return nil, err
	}

	return frontmatter, nil
}
