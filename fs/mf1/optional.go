package mf1

import (
	"errors"
	"os"
)

//LastStater is optional, stat  file
type LastStater interface {
	LastStaterTry(name string) (os.FileInfo, bool, error)
}

// SymblicLink is optional,
type SymblicLink interface {
	LastStater
	Linker
	LinkReader
}

type Linker interface {
	SymblicLinkTry(oldName string, newname string) error
}
type LinkReader interface {
	ReadLinkTry(name string) (string, error)
}

var ErrNoLastStater = errors.New("last stat not supported")
var ErrNoSymblicLink = errors.New("symblic link not supported")
var ErrNoReadLink = errors.New("readlink not supported")
