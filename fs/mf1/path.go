package mf1

import (
	"os"
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
	names, err := file.Readdirnames(-1)
	file.Close()
	if err != nil {
		return nil, err
	}
	sort.Strings(names)
	return names, nil
}

// walk recursively descends path, calling walkFn
// adapted from https://golang.org/src/path/filepath/path.go
func walk(mf Mf, path string, info os.FileInfo, walkFunc filepath.WalkFunc) error {

	err := walkFunc(path, info, nil)
	if err != nil {
		if info.IsDir() && err == filepath.SkipDir {
			return nil
		}
		return err
	}

	if !info.IsDir() {
		return nil
	}
	names, err := readDirNames(mf, path)
	if err != nil {
		return walkFunc(path, info, err)
	}
	for _, name := range names {
		filename := filepath.Join(path, name)
		fileInfo, err := lstatTry(mf, filename)
		if err != nil {
			if err := walkFunc(filename, info, err); err != nil && err != filepath.SkipDir {
				return err
			}
		} else {
			err = walk(mf, filename, info, walkFunc)
			if err != nil {
				if !fileInfo.IsDir() || err != filepath.SkipDir {
					return err
				}
			}
		}

	}
	return nil
}

func lstatTry(mf Mf, path string) (os.FileInfo, error) {
	if lmf, ok := mf.(Lstater); ok {
		fileInfo, _, err := lmf.LstaterTry(path)
		return fileInfo, err
	}
	return mf.Stat(path)
}

func Walk(mf Mf, root string, walkFunc filepath.WalkFunc) error {
	info, err := lstatTry(mf, root)
	if err != nil {
		return walkFunc(root, nil, err)
	}
	return walk(mf, root, info, walkFunc)
}
