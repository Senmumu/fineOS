//Optional contains optional interface in file system
package mf1

import (
	"errors"
	"os"
)

//LastStater is optional, stat  file
type Lstater interface {
	LstaterTry(name string) (os.FileInfo, bool, error)
}

// SymblicLink is optional,
type SymblicLink interface {
	Lstater
	Linker
	LinkReader
}

type Linker interface {
	SymblicLinkTry(oldName string, newname string) error
}
type LinkReader interface {
	ReadLinkTry(name string) (string, error)
}

var ErrNoLstater = errors.New("lstat not supported")
var ErrNoSymblicLink = errors.New("symblic link not supported")
var ErrNoReadLink = errors.New("readlink not supported")
