package toolkit

import (
	"errors"
	"os"
	"regexp"
	"strings"
)

// CreateDirIfNotExist creates a directory and its parents if it does not exist
func (t *Tools) CreateDirIfNotExist(path string) error {
	const mode = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, mode)
		if err != nil {
			return err
		}

	}

	return nil
}

// Slugify converts a string to a slug
func (t *Tools) Slugify(s string) (string, error) {
	if s == "" {
		return "", errors.New("empty string")
	}

	var re = regexp.MustCompile("[^a-z0-9]+")
	slug := strings.Trim(re.ReplaceAllString(strings.ToLower(s), "-"), "-")

	if len(slug) == 0 {
		return "", errors.New("empty slug is generated")
	}

	return slug, nil
}
