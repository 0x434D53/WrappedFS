package wrappedfs

import (
	"io"
	"net/http"
	"os"
	"path"
	"testing"
)

type mockFS struct {
	path string
}

func (mfs *mockFS) Open(name string) (http.File, error) {
	if mfs.path == path.Clean(name) {
		return &mockFile{}, nil
	}

	return nil, os.ErrNotExist
}

type mockFile struct {
	os.FileInfo
	io.Closer
	io.Reader
}

func (f *mockFile) Seek(offset int64, whence int) (ret int64, err error) {
	return 0, nil
}

func (f *mockFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *mockFile) Readdir(count int) ([]os.FileInfo, error) {
	return make([]os.FileInfo, 0), nil
}

func (f *mockFile) Read(p []byte) (n int, err error) {
	return 0, nil
}

func TestMockFS(t *testing.T) {
	fs := mockFS{path: "/p1/p2"}

	_, err := fs.Open("/p1/p2")

	if err != nil {
		t.Fatalf("File should be found, but wasn't")
	}

	_, err = fs.Open("/p1//p2")

	if err != nil {
		t.Fatalf("File with Double Slashes should also be found")
	}

	f, err := fs.Open("/p3")

	if err != os.ErrNotExist {
		t.Fatalf("os.ErrNotExist was expected but not given")
	}

	if f != nil {
		t.Fatalf("In case of error the returned file should be nil")
	}
}

func TestWrappedFS(t *testing.T) {
	fs := &mockFS{path: "/p1/p2"}
	wfs := NewWrappedFS(fs, "/p1")

	_, err := wfs.Open("/p2")

	if err != nil {
		t.Fatalf("/p2 should be found. Wasn't")
	}

	_, err = wfs.Open("//p2")

	if err != nil {
		t.Fatalf("//p2 shoud be found. Wasn't")
	}
}

func TestEscaping(t *testing.T) {
	fs := &mockFS{path: "/p1/p2"}
	wfs := NewWrappedFS(fs, "/p3")

	_, err := wfs.Open("../p1/p2")

	if err != os.ErrNotExist {
		t.Fatalf("The Path escaped the Prefix. Should NOT be possible")
	}
}
