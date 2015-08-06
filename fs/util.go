package fs

import (
	"strings"

	"github.com/scottferg/Dropbox-Go/dropbox"
)

func MakeURI(path string) dropbox.Uri {
	return dropbox.Uri{
		Root: dropbox.RootDropbox,
		Path: path,
	}
}

func NameFromPath(path string) (name string) {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}
