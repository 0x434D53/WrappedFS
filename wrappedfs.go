package wrappedfs

import "net/http"

type WrappedFS struct {
	fs  http.FileSystem
	dir string
}

func (wfs *WrappedFS) Open(name string) (http.File, error) {
	return nil, nil
}

func NewWrappedFS(fs http.FileSystem, path string) (http.FileSystem, error) {
	return nil, nil
}
