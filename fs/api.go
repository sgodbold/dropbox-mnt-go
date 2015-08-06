package fs

import (
	"github.com/hanwen/go-fuse/fuse/pathfs"
	"github.com/scottferg/Dropbox-Go/dropbox"
)

// Globals
var Config Configuration
var Session dropbox.Session

type Configuration struct {
	AppKey      string
	AppSecret   string
	AccessType  string
	TokenSecret string
	TokenKey    string
}

type DropboxFs struct {
	pathfs.FileSystem
}

type Metadata struct {
	Hash     string
	Bytes    int
	Modified string
	Path     string
	IsDir    bool
	Contents []string // list of paths
}
