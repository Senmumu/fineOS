package mf1

import (
	"os"
	"time"
)

var _ Lstater = (*MfOs)(nil)

func NewMfOs() Mf {
	return &MfOs{}
}

type MfOs struct {
}

func (MfOs) Name() string {
	return "MfOs"
}

func (MfOs) Create(name string) (MfFile, error) {
	file, err := os.Create(name)
	if file == nil {
		return nil, err
	}
	return file, err
}
func (MfOs) Mkdir(name string, mode os.FileMode) error {
	return os.Mkdir(name, mode)
}

func (MfOs) MkdirAll(path string, mode os.FileMode) error {
	return os.MkdirAll(path, mode)
}

func (MfOs) Open(name string) (MfFile, error) {
	file, err := os.Open(name)
	if file == nil {
		return nil, err
	}
	return file, err
}

func (MfOs) OpenFile(name string, flag int, mode os.FileMode) (MfFile, error) {
	file, err := os.OpenFile(name, flag, mode)
	if file == nil {
		return nil, err
	}
	return file, err
}

func (MfOs) Remove(name string) error {
	return os.Remove(name)
}
func (MfOs) RemoveAll(name string) error {
	return os.RemoveAll(name)
}

func (MfOs) Rename(oldName, newName string) error {
	return os.Rename(oldName, newName)
}

func (MfOs) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}
func (MfOs) Chown(name string, userId, groupId int) error {
	return os.Chown(name, userId, groupId)
}

func (MfOs) Chmod(name string, mode os.FileMode) error {
	return os.Chmod(name, mode)
}

func (MfOs) Chtimes(name string, appendTime time.Time, modifyTime time.Time) error {
	return os.Chtimes(name, appendTime, modifyTime)
}

func (MfOs) LstatTry(name string) (os.FileInfo, bool, error) {
	fileInfo, err := os.Lstat(name)
	return fileInfo, true, err
}

func (MfOs) SymbliclinkTry(oldName, newName string) error {
	return os.Symlink(oldName, newName)
}

func (MfOs) ReadLinkTry(name string) (string, error) {
	return os.Readlink(name)
}
