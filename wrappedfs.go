package wrappedfs

import (
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"strings"
)

// WrappedFS wrapped an http.FileSystem, but prepends the initalized dir to the path
type WrappedFS struct {
	fs  http.FileSystem
	dir string
}

// Open opens a file from the unterlying http.Filesystem at wfs.dir/name
func (wfs *WrappedFS) Open(name string) (http.File, error) {
	if filepath.Separator != '/' && strings.IndexRune(name, filepath.Separator) >= 0 || strings.Contains(name, "\x00") {
		return nil, fmt.Errorf("Invalid character in file path")
	}

	dir := wfs.dir
	path := filepath.Join(dir, filepath.FromSlash(path.Clean("/"+name)))

	return wfs.fs.Open(path)
}

// New creates a new WrappedFS with the given dir
func New(fs http.FileSystem, dir string) http.FileSystem {
	wfs := WrappedFS{fs: fs, dir: path.Clean(dir)}

	return &wfs
}
