package wrappedfs

import (
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"strings"
)

type WrappedFS struct {
	fs  http.FileSystem
	dir string
}

func (wfs *WrappedFS) Open(name string) (http.File, error) {
	if filepath.Separator != '/' && strings.IndexRune(name, filepath.Separator) >= 0 || strings.Contains(name, "\x00") {
		return nil, fmt.Errorf("Invalid character in file path")
	}

	dir := wfs.dir
	path := filepath.Join(dir, filepath.FromSlash(path.Clean("/"+name)))

	return wfs.fs.Open(path)
}

func NewWrappedFS(fs http.FileSystem, dir string) http.FileSystem {
	wfs := WrappedFS{fs: fs, dir: path.Clean(dir)}

	return &wfs
}
