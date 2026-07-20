package scan

import (
	"os"
	"path/filepath"
	"strings"
)

func ScanDirectory(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files []string

	for _, entry := range entries {		
		path := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			subFiles, err := ScanDirectory(path)
			if err != nil {
				return nil, err
			}
			files = append(files, subFiles...)
		} else {
			if strings.HasSuffix(entry.Name(), ".go") {
    			// It's a Go source file.
				files = append(files, path)
			}
		}
	}

	return files, nil
}