package web

import (
	"net/http"
	"os"

	"github.com/spy16/droplets/pkg/logger"
)

func newSafeFileSystemServer(lg logger.Logger, root string) http.Handler {
	sfs := &safeFileSystem{
		fs:     http.Dir(root),
		Logger: lg,
	}
	return http.FileServer(sfs)
}

// safeFileSystem implements http.FileSystem. It is used to prevent directory
// listing of static assets.
type safeFileSystem struct {
	logger.Logger

	fs http.FileSystem
}

func (sfs safeFileSystem) Open(path string) (http.File, error) {
	f, err := sfs.fs.Open(path)
	if err != nil {
		sfs.Warnf("failed to open file '%s': %v", path, err)
		return nil, err
	}

	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		sfs.Warnf("path '%s' is a directory, rejecting static path request", path)
		return nil, os.ErrNotExist
	}

	return f, nil
}
