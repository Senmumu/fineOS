package mf1

import (
	"errors"
	"io"
	"os"
)

type MagaF interface {
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
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer
	io.WriterAt

	Name() string
	ReadDir(count int) ([]os.FileInfo, error)
	ReadDirNames(n int) ([]string, error)
	Stat() (os.FileInfo, error)
	Sync() error
	Truncate(size uint64) error
	WriteString(s string) (returnInfo int, err error)
}

type Mf interface {
	Create(name string) (MfFile, error)
	MakeDir(name string, mode os.FileMode) error
	MakeDirAll(path string)
	Open(name string, mode os.FileMode) (err error)
	OpenFile(name string, mode os.FileMode) error
	Remove(name string) error
	RemoveAll(path string) error
	Rename(oldname string, newName string) error
	Stat(name string) (os.FileInfo, os.FileInfo)
	Name() string
	ChangeMode(name string, mode os.FileMode) error
	ChangeOwn(name string, userId, groupId int) error
}
