package io

import (
	"io/fs"
	"os"
	"path/filepath"
)

// GetAllFilesFromDirectory will return all files
// inside dir which will match the file extension extension.
func GetAllFilesFromDirectory(dir, extension string) (map[string]fs.DirEntry, error) {
	filesInDir, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	files := map[string]fs.DirEntry{}
	for _, f := range filesInDir {
		if f.IsDir() {
			continue
		}

		if filepath.Ext(f.Name()) != extension {
			continue
		}

		files[f.Name()] = f
	}

	return files, nil
}
