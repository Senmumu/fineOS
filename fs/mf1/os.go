package mf1

import (
	"os"
	"time"
)

var _ LastStater = (*MfOs)(nil)

func NewMfOs() Mf {
	return &MfOs{}
}

type MfOs struct {
}

func (MfOs) Name() string {
	return "MfOs"
}

func (MfOs) Create(name string) (MfFile, error) {
	file, e := os.Create(name)
	if file == nil {
		return nil, e
	}
	return file, e
}
func MakeDir(name string, permition os.FileMode) error {
	return os.Mkdir(name, permition)
}

func (MfOs) MakeDirAll(path string, permition os.FileMode) error {
	return os.MkdirAll(path, permition)
}

func (MfOs) Open(name string) (MfFile, error) {
	file, err := os.Open(name)
	if file == nil {
		return nil, err
	}
	return file, err
}

func OpenFile(name string, flag int, permition os.FileMode) (File, error) {
	file, err := os.OpenFile(name, flag, permition)
	if file == nil {
		return nil, err
	}
	return file, err
}

func (MfOs) Remove(name string) error {
	return os.Remove(name)
}

func (MfOs) Rename(oldName, newName string) error {
	return os.Rename(oldName, newName)
}

func (MfOs) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}
func (MfOs) ChangeOwn(name string, userId, groupId int) error {
	return os.Chown(name, userId, groupId)
}

func (MfOs) ChangeMode(name string, mode os.FileMode) error {
	return os.Chmod(name, mode)
}

func (MfOs) ChangeTimes(name string, appendTime time.Time, modifyTime time.Time) error {
	return os.Chtimes(name, appendTime, modifyTime)
}

func (MfOs) LastStatIfPossible(name string) (os.FileInfo, bool, error) {
	fileInfo, err := os.Lstat(name)
	return fileInfo, true, err
}

func (MfOs) SymbliclinkIfPossible(oldName, newName string) error {
	return os.Symlink(oldName, newName)
}

func (MfOs) ReadLinkIfPossible(name string) (string, error) {
	return os.Readlink(name)
}
