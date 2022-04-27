package mf1

import (
	"regexp"
	"syscall"
)

type ReMf struct {
	re     *regexp.Regexp
	source Mf
}

func NewReMf(source Mf, re *regexp.Regexp) Mf {
	return &ReMf{source: source, re: re}
}

type ReFile struct {
	f  MfFile
	re *regexp.Regexp
}

func (reMf *ReMf) matchName(name string) error {
	if reMf.re == nil {
		return nil
	}
	if reMf.re.MatchString(name) {
		return nil
	}
	return syscall.ENOENT
}

func (reMf *ReMf) dirOrMatches(name string) error {
	dir, err := IsDir(reMf.source, name)
	if err != nil {
		return err
	}
	if dir {
		return nil
	}
	return reMf.matchName(name)
}
