package wrappedfs

import (
	"net/http"
	"path"
)

type WrappedFS struct {
	fs  http.FileSystem
	dir string
}

func (wfs *WrappedFS) Open(name string) (http.File, error) {
	name = path.Clean(name)
	name = path.Join(wfs.dir, name)

	return wfs.fs.Open(name)
}

func NewWrappedFS(fs http.FileSystem, dir string) http.FileSystem {
	wfs := WrappedFS{fs: fs, dir: path.Clean(dir)}

	return &wfs
}
