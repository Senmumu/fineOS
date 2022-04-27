package mf1

import (
	"path/filepath"
	"sort"
)

// readDirNames reads the directory named by dirname and returns
// a sorted list of directory entries.
// adapted from https://golang.org/src/path/filepath/path.go
func readDirNames(mf Mf, dirName string) ([]string, error) {
	file, err := mf.Open(dirName)
	if err != nil {
		return nil, err
	}
	names, err := file.ReadDirNames(-1)
	file.Close()
	if err != nil {
		return nil, err
	}
	sort.Strings(names)
	return names, nil
}

// walk recursively descends path, calling walkFn
// adapted from https://golang.org/src/path/filepath/path.go
func walk(mf Mf, path string, walkFn filepath.WalkFunc) {

	err := walkFn(path, info, nil)
	if err != nil {
		if info.IsDir() && err == filePath.SkipDir {
			return nil
		}
		return err
	}

	if !info.IsDir() {
		return nil
	}
	name, err := readDirNames(fs, path)
	if err != nil {
		return walkFn(path, info, err)
	}
	for _, name := range names {
		filename := filepath.Join(path, name)
		fileInfo, err := lastStatTry(fs, filename)

	}
}
