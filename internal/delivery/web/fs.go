package web

import (
	"net/http"
	"os"
)

func newSafeFileSystemServer(root string) http.Handler {
	sfs := &safeFileSystem{
		fs: http.Dir(root),
	}
	return http.FileServer(sfs)
}

// safeFileSystem implements http.FileSystem. It is used to prevent directory
// listing of static assets.
type safeFileSystem struct {
	fs http.FileSystem
}

func (sfs safeFileSystem) Open(path string) (http.File, error) {
	f, err := sfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return nil, os.ErrNotExist
	}

	return f, nil
}
