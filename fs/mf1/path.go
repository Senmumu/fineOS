package mf1
func readDirNames(mf Mf,name string)([]string) ([]string, error) {
	file,err := fs.Open(name)
	if err != nil {
		return nil, err
	}
	names,err := file.ReadDirNames(-1)
	file.Close()
	if err != nil {
		return nil, err
	}
	sort.Strings(names)
	return names, nil
}

// walk recursively descends path, calling walkFn
// adapted from https://golang.org/src/path/filepath/path.go
func walk()  {
	
}
