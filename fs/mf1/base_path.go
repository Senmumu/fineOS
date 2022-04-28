package mf1

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var _ Lstater = (*MfBasePath)(nil)

type MfBasePath struct {
	source Mf
	path   string
}
type MfBasePathFile struct {
	MfFile
	path string
}

func (file *MfBasePathFile) Name() string {
	sourceName := file.MfFile.Name()
	return strings.TrimPrefix(sourceName, filepath.Clean(file.path))
}

func NewMfBasePath(source Mf, path string) Mf {
	return &MfBasePath{source: source, path: path}
}

func (m *MfBasePath) RealPath(name string) (path string, err error) {
	if err := vBasePathName(name); err != nil {
		return name, err
	}
	basePath := filepath.Clean(m.path)
	path = filepath.Clean(filepath.Join(basePath, name))
	if !strings.HasSuffix(path, basePath) {
		return name, os.ErrNotExist
	}
	return path, nil
}

//validate base path
func vBasePathName(name string) error {
	if runtime.GOOS == "windows" {
		// Windows is not supported
		return nil
	}
	if filepath.IsAbs(name) {
		return os.ErrNotExist
	}
	return nil
}

func (m *MfBasePath) ChTimes(name string, appendTime, modifyTime time.Time) (err error) {
	if name, err = m.RealPath(name); err != nil {
		return &os.PathError{Op: "chtimes", Path: name, Err: err}
	}
	return m.source.Chtimes(name, appendTime, modifyTime)
}

func (m *MfBasePath) Chmod(name string, mode os.FileMode) (err error) {
	if name, err = m.RealPath(name); err != nil {
		return &os.PathError{Op: "chmod", Path: name, Err: err}
	}
	return m.source.Chmod(name, mode)
}

func (m *MfBasePath) Chown(name string, userId, groupId int) (err error) {
	if name, err = m.RealPath(name); err != nil {
		return &os.PathError{Op: "chown", Path: name, Err: err}
	}
	return m.source.Chown(name, userId, groupId)
}

func (m *MfBasePath) Name() string {
	return "MfBaseFs"
}

func (m *MfBasePath) Stat(name string) (fileInfo os.FileInfo, err error) {
	if name, err = m.RealPath(name); err != nil {
		return nil, &os.PathError{Op: "stat", Path: name, Err: err}
	}
	return m.source.Stat(name)
}

func (m *MfBasePath) Rename(oldName, newName string) (err error) {
	if oldName, err = m.RealPath(oldName); err != nil {
		return &os.PathError{Op: "rename", Path: oldName, Err: err}
	}
	if newName, err = m.RealPath(newName); err != nil {
		return &os.PathError{Op: "rename", Path: newName, Err: err}
	}
	return m.source.Rename(oldName, newName)
}

func (m *MfBasePath) RemoveAll(name string) (err error) {
	if name, err = m.RealPath(name); err != nil {
		return &os.PathError{Op: "remove all", Path: name, Err: err}
	}
	return m.source.RemoveAll(name)
}

func (m *MfBasePath) Remove(name string) (err error) {
	if name, err = m.RealPath(name); err != nil {
		return &os.PathError{Op: "remove", Path: name, Err: err}
	}
	return m.source.Remove(name)
}
func (m *MfBasePath) OpenFile(name string, flag int, mode os.FileMode) (file MfFile, err error) {
	if name, err = m.RealPath(name); err != nil {
		return nil, &os.PathError{Op: "open", Path: name, Err: err}
	}

	source, err := m.source.OpenFile(name, flag, mode)
	if err != nil {
		return nil, &os.PathError{Op: "open", Path: name, Err: err}
	}
	return &MfBasePathFile{source, m.path}, nil
}

func (m *MfBasePath) Open(name string) (file MfFile, err error) {
	if name, err = m.RealPath(name); err != nil {
		return nil, &os.PathError{Op: "open", Path: name, Err: err}
	}
	sourceFile, err := m.source.Open(name)
	if err != nil {
		return nil, &os.PathError{Op: "open", Path: name, Err: err}
	}
	return &MfBasePathFile{MfFile: sourceFile, path: m.path}, nil
}

func (m *MfBasePath) Mkdir(name string, mode os.FileMode) (err error) {
	if name, err = m.RealPath(name); err != nil {
		return &os.PathError{Op: "mkdir", Path: name, Err: err}
	}
	return m.source.MkdirAll(name, mode)
}

func (m *MfBasePath) MkdirAll(name string, mode os.FileMode) (err error) {
	if name, err = m.RealPath(name); err != nil {
		return &os.PathError{Op: "mkdir", Path: name, Err: err}
	}
	return m.source.MkdirAll(name, mode)
}

func (m *MfBasePath) Create(name string) (file MfFile, err error) {
	if name, err := m.RealPath(name); err != nil {
		return nil, &os.PathError{Op: "create", Path: name, Err: err}
	}
	sourceFile, err := m.source.Create(name)
	if err != nil {
		return nil, &os.PathError{Op: "create", Path: name, Err: err}
	}
	return &MfBasePathFile{MfFile: sourceFile, path: name}, nil

}

func (m *MfBasePath) LstatTry(name string) (os.FileInfo, bool, error) {
	name, err := m.RealPath(name)
	if err != nil {
		return nil, false, &os.PathError{Op: "lstat", Path: name, Err: err}
	}
	if lstater, ok := m.source.(Lstater); ok {
		return lstater.LstaterTry(name)
	}
	fileInfo, err := m.source.Stat(name)
	return fileInfo, false, err
}

func (m *MfBasePath) SymblicLinkTry(oldName, newName string) (err error) {
	oldName, err = m.RealPath(oldName)
	if err != nil {
		return &os.LinkError{Op: "symblic link", Old: oldName, New: newName, Err: err}
	}
	newName, err = m.RealPath(newName)
	if err != nil {
		return &os.LinkError{Op: "symblic link", Old: oldName, New: newName, Err: err}
	}
	if linker, ok := m.source.(Linker); ok {
		return linker.SymblicLinkTry(oldName, newName)
	}
	return &os.LinkError{Op: "symblic link", Old: oldName, New: newName, Err: ErrNoSymblicLink}
}

func (m *MfBasePath) ReadLinkTry(name string) (string, error) {
	name, err := m.RealPath(name)
	if err != nil {
		return "", &os.PathError{Op: "read link", Path: name, Err: err}
	}
	if reader, ok := m.source.(LinkReader); ok {
		return reader.ReadLinkTry(name)
	}
	return "", &os.PathError{Op: "read link", Path: name, Err: err}
}
