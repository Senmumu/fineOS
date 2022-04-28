package mf1

import (
	"errors"
	"io"
	"os"
	"time"
)

type MagaF struct {
	Mf
}

var (
	ErrFileClosed   = errors.New("File is closed")
	ErrOutOfRange   = errors.New("Out of range")
	ErrTooLarge     = errors.New("Too large")
	ErrFileNotFound = errors.New("File not found")
	ErrFileExists   = errors.New("File does not exist")
)

type MfFile interface {
	io.Seeker
	io.Reader
	io.ReaderAt
	io.Writer
	io.WriterAt
	io.Closer

	Name() string
	Readdir(count int) ([]os.FileInfo, error)
	Readdirnames(n int) ([]string, error)
	Stat() (os.FileInfo, error)
	Sync() error
	Truncate(size int64) error
	WriteString(s string) (returnInfo int, err error)
}

type Mf interface {
	Create(name string) (MfFile, error)
	Mkdir(name string, mode os.FileMode) error
	MkdirAll(path string, mode os.FileMode) error
	Open(name string) (MfFile, error)
	OpenFile(name string, flag int, mode os.FileMode) (MfFile, error)
	Remove(name string) error
	RemoveAll(path string) error
	Rename(oldname string, newName string) error
	Stat(name string) (os.FileInfo, error)
	Name() string
	Chmod(name string, mode os.FileMode) error
	Chown(name string, userId, groupId int) error
	Chtimes(name string, appendTime time.Time, modifyTime time.Time) error
}
