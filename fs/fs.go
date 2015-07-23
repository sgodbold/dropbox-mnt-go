package fs

import (
	"encoding/json"
	"log"
	"os"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
)

var Config Configuration

type DropboxFs struct {
	pathfs.FileSystem
}

type Configuration struct {
	AppKey      string
	AppSecret   string
	AccessType  string
	TokenSecret string
	TokenKey    string
}

// LoadConfig loads 'filename' into a global Configuration struct.
func LoadConfig(filename string) (err error) {
	// Clear any current values
	Config = Configuration{}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}
	return nil
}

// MountFs mounts the filesystem at 'path'. Loads config and starts dropbox session.
func MountFs(path string) {
	err := LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Config fail: %v\n", err)
	}
	LoadSession()
	nfs := pathfs.NewPathNodeFs(&DropboxFs{FileSystem: pathfs.NewDefaultFileSystem()}, nil)
	server, _, err := nodefs.MountRoot(path, nfs.Root(), nil)
	if err != nil {
		log.Fatalf("Mount fail: %v\n", err)
	}
	server.Serve()
}

func (me *DropboxFs) GetAttr(name string, context *fuse.Context) (*fuse.Attr, fuse.Status) {
	attr := &fuse.Attr{}
	if name != "" {
		attr.Mode = fuse.S_IFREG | 0644
		attr.Size = uint64(len(name))
	} else {
		attr.Mode = fuse.S_IFDIR | 0755
	}
	return attr, fuse.OK
}

func (me *DropboxFs) OpenDir(name string, context *fuse.Context) (c []fuse.DirEntry, code fuse.Status) {
	log.Println("OpenDir")
	if name == "" {
		entry := fuse.DirEntry{}
		file_names, err := FilenamesInDir(name)
		if err != nil {
			return nil, fuse.ENOENT
		}
		for _, file := range file_names {
			entry.Name = file
			entry.Mode = fuse.S_IFREG
			c = append(c, entry)
		}
		return c, fuse.OK
	}
	return nil, fuse.ENOENT
}
