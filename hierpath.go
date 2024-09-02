package hierpath

import (
	"bytes"
	"os"
	"path/filepath"
	"strconv"
)

func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func MakeDir(dir string) (e error) {
	if !IsExists(dir) {
		e = os.Mkdir(dir, 0755) //os.ModeDir
	}
	return
}

func MakeHierarchicalPath(levels int, desDir, filename string) (fullpath string) {
	fpath := desDir
	ss := filename
	b := []byte(ss)
	i := bytes.LastIndex(b, []byte("."))
	if i > 0 {
		ss = string(b[:i])
	}
	b = []byte(ss)
	n := len(b)
	for i = 0; i < levels; i++ {
		if i < n {
			fpath = filepath.Join(fpath, strconv.Itoa(int(b[n-1-i])))
			MakeDir(fpath)
		} else {
			break
		}
	}
	fullpath = filepath.Join(fpath, filename)
	return
}

func MappingHierarchicalPath(levels int, desDir, filename string) (fullpath string) {
	fpath := desDir
	ss := filename
	b := []byte(ss)
	i := bytes.LastIndex(b, []byte("."))
	if i > 0 {
		ss = string(b[:i])
	}
	b = []byte(ss)
	n := len(b)
	for i = 0; i < levels; i++ {
		if i < n {
			fpath = filepath.Join(fpath, strconv.Itoa(int(b[n-1-i])))
		} else {
			break
		}
	}
	fullpath = filepath.Join(fpath, filename)
	return
}

func UrlHierarchicalPath(levels int, filename string) (urlpath string) {
	urlpath = ""
	ss := filename
	b := []byte(ss)
	i := bytes.LastIndex(b, []byte("."))
	if i > 0 {
		ss = string(b[:i])
	}
	b = []byte(ss)
	n := len(b)
	for i = 0; i < levels; i++ {
		if i < n {
			urlpath += strconv.Itoa(int(b[n-1-i])) + "/"
		} else {
			break
		}
	}
	urlpath += filename
	return
}
