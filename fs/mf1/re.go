// re provide a sub set of regexp file methods
package mf1

import (
	"fmt"
	"os"
	"regexp"
	"syscall"
	"time"
)

type ReMf struct {
	re     *regexp.Regexp
	source Mf
}

func NewReMf(source Mf, re *regexp.Regexp) Mf {
	return &ReMf{source: source, re: re}
}

type ReFile struct {
	mfFile MfFile
	re     *regexp.Regexp
}

func (r *ReMf) Name() string {
	return "ReMf"
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

func (reMf *ReMf) dirOrMatch(name string) error {
	dir, err := IsDir(reMf.source, name)
	if err != nil {
		return err
	}
	if dir {
		return nil
	}
	return reMf.matchName(name)
}

func (r *ReMf) Open(name string) (MfFile, error) {
	dir, err := IsDir(r.source, name)
	if err != nil {
		return nil, err
	}
	if !dir {
		if err := r.matchName(name); err != nil {
			return nil, err
		}
	}
	file, err := r.source.Open(name)
	if err != nil {
		return nil, err
	}
	return &ReFile{mfFile: file, re: r.re}, nil
}

func (r *ReMf) Mkdir(name string, mode os.FileMode) error {
	return r.source.Mkdir(name, mode)
}

func (r *ReMf) MkdirAll(name string, mode os.FileMode) error {
	return r.source.MkdirAll(name, mode)
}
func (r *ReMf) Create(name string) (MfFile, error) {
	if err := r.matchName(name); err != nil {
		return nil, err
	}
	return r.source.Create(name)
}

func (file *ReFile) Close() error {
	return file.mfFile.Close()
}

func (file *ReFile) Read(s []byte) (int, error) {
	return file.mfFile.Read(s)
}
func (file *ReFile) ReadAt(s []byte, offset int64) (int, error) {
	return file.mfFile.ReadAt(s, offset)
}

func (file *ReFile) Write(s []byte) (int, error) {
	return file.mfFile.Write(s)
}

func (file *ReFile) WriteAt(s []byte, offset int64) (int, error) {
	return file.mfFile.WriteAt(s, offset)
}
func (file *ReFile) Name() string {
	return file.mfFile.Name()
}

func (file *ReFile) Readdir(n int) (fileInfo []os.FileInfo, err error) {
	var reFileInfo []os.FileInfo
	reFileInfo, err = file.mfFile.Readdir(n)
	if err != nil {
		return nil, err
	}
	for _, i := range reFileInfo {
		if i.IsDir() || file.re.MatchString(i.Name()) {
			fileInfo = append(fileInfo, i)
		}
	}
	return fileInfo, nil
}

func (file *ReFile) Readdirnames(n int) (names []string, err error) {
	fileInfo, err := file.Readdir(n)
	if err != nil {
		return nil, err
	}
	for _, i := range fileInfo {
		names = append(names, i.Name())
	}
	return names, nil
}

func (file *ReFile) Seek(offset int64, whence int) (int64, error) {
	return file.mfFile.Seek(offset, whence)
}

func (file *ReFile) Stat() (os.FileInfo, error) {
	return file.mfFile.Stat()
}
func (file *ReFile) Sync() error {
	return file.mfFile.Sync()
}
func (file *ReFile) Truncate(s int64) error {
	return file.mfFile.Truncate(s)
}
func (file *ReFile) WriteString(s string) (n int, err error) {
	return file.mfFile.WriteString(s)
}

func (r *ReMf) Chtimes(name string, appendTime, modifyTime time.Time) error {
	if err := r.dirOrMatch(name); err != nil {
		return err
	}
	return r.source.Chtimes(name, appendTime, modifyTime)
}

func (r *ReMf) Chmod(name string, mode os.FileMode) error {
	if err := r.dirOrMatch(name); err != nil {
		return err
	}
	return r.source.Chmod(name, mode)
}

func (r *ReMf) Chown(name string, uid int, gid int) error {
	if err := r.dirOrMatch(name); err != nil {
		return err
	}
	return r.source.Chown(name, uid, gid)
}

func (r *ReMf) Stat(name string) (os.FileInfo, error) {
	if err := r.dirOrMatch(name); err != nil {
		return nil, err
	}
	return r.source.Stat(name)
}

func (r *ReMf) Rename(oldName, newName string) error {
	dir, err := IsDir(r.source, oldName)
	if err != nil {
		return err
	}
	if dir {
		return nil
	}
	if err := r.matchName(oldName); err != nil {
		return err
	}
	if err := r.matchName(newName); err != nil {
		return err
	}
	fmt.Println("modify success")
	return r.source.Rename(oldName, newName)

}

func (r *ReMf) RemoveAll(path string) error {
	dir, err := IsDir(r.source, path)
	if err != nil {
		return err
	}
	if !dir {
		if err := r.matchName(path); err != nil {
			return err
		}
	}
	return r.source.RemoveAll(path)
}

func (r *ReMf) Remove(name string) error {
	if err := r.dirOrMatch(name); err != nil {
		return err
	}
	return r.source.Remove(name)
}

func (r *ReMf) OpenFile(name string, flag int, mode os.FileMode) (MfFile, error) {
	if err := r.dirOrMatch(name); err != nil {
		return nil, err
	}
	return r.source.OpenFile(name, flag, mode)
}
