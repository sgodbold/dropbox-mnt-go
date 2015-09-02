package fs

import (
	"log"
	"time"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

func get_path(name string) string {
	return Config.DropBoxMnt + "/" + name
}

func (me *DropboxFs) GetAttr(name string, context *fuse.Context) (*fuse.Attr, fuse.Status) {
	path := get_path(name)
	log.Printf("GetAttr: '%s'\n", name)
	attr := fuse.Attr{}
	// XXX: handle this error
	data, _ := Cache.Metadata.Get(path)
	if data.IsDir {
		attr.Mode = fuse.S_IFDIR | 0755
	} else {
		attr.Mode = fuse.S_IFREG | 0644
		attr.Size = uint64(len(path))
	}
	return &attr, fuse.OK
}

func (me *DropboxFs) Chmod(name string, mode uint32, context *fuse.Context) (code fuse.Status) {
	log.Printf("Chmod: '%s'\n", name)
	return fuse.ENOSYS
}

func (me *DropboxFs) Chown(name string, uid uint32, gid uint32, context *fuse.Context) (code fuse.Status) {
	log.Printf("Chown: '%s'\n", name)
	return fuse.ENOSYS
}

// unsupported
func (me *DropboxFs) Utimens(name string, Atime *time.Time, Mtime *time.Time, context *fuse.Context) (code fuse.Status) {
	log.Printf("Utimens: '%s'\n", name)
	return fuse.ENOSYS
}

func (me *DropboxFs) Truncate(name string, size uint64, context *fuse.Context) (code fuse.Status) {
	log.Printf("Truncate: '%s'\n", name)
	return fuse.ENOSYS
}

// unsupported
func (me *DropboxFs) Access(name string, mode uint32, context *fuse.Context) (code fuse.Status) {
	log.Printf("Access: '%s'\n", name)
	return fuse.ENOSYS
}

func (me *DropboxFs) Link(oldName string, newName string, context *fuse.Context) (code fuse.Status) {
	log.Printf("Link: '%s'\n", oldName)
	return fuse.ENOSYS
}

func (me *DropboxFs) Mkdir(name string, mode uint32, context *fuse.Context) fuse.Status {
	log.Printf("Mkdir: '%s'\n", name)
	return fuse.ENOSYS
}

func (me *DropboxFs) Mknod(name string, mode uint32, dev uint32, context *fuse.Context) fuse.Status {
	log.Printf("Mknod: '%s'\n", name)
	return fuse.ENOSYS
}

func (me *DropboxFs) Rename(oldName string, newName string, context *fuse.Context) (code fuse.Status) {
	log.Printf("Rename: '%s'\n", oldName)
	return fuse.ENOSYS
}

func (me *DropboxFs) Rmdir(name string, context *fuse.Context) (code fuse.Status) {
	log.Printf("Rmdir: '%s'\n", name)
	return fuse.ENOSYS
}

func (me *DropboxFs) Unlink(name string, context *fuse.Context) (code fuse.Status) {
	log.Printf("Unlink: '%s'\n", name)
	return fuse.ENOSYS
}

func (me *DropboxFs) GetXAttr(name string, attribute string, context *fuse.Context) (data []byte, code fuse.Status) {
	log.Printf("GetXAttr: '%s'\n", name)
	return nil, fuse.ENOSYS
}

func (me *DropboxFs) ListXAttr(name string, context *fuse.Context) (attributes []string, code fuse.Status) {
	log.Printf("ListXAttr: '%s'\n", name)
	return nil, fuse.ENOSYS
}

func (me *DropboxFs) RemoveXAttr(name string, attr string, context *fuse.Context) fuse.Status {
	log.Printf("RemoveXAttr: '%s'\n", name)
	return fuse.ENOSYS
}

func (me *DropboxFs) SetXAttr(name string, attr string, data []byte, flags int, context *fuse.Context) fuse.Status {
	log.Printf("SetXAttr: '%s'\n", name)
	return fuse.ENOSYS
}

func (me *DropboxFs) Open(name string, flags uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
	path := get_path(name)
	log.Printf("Open: '%s'\n", name)
	if flags&fuse.O_ANYWRITE != 0 {
		return nil, fuse.EPERM
	}
	// XXX: errors!
	res, _ := GetFile(path)
	return nodefs.NewDataFile(res), fuse.OK
}

func (me *DropboxFs) Create(name string, flags uint32, mode uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
	log.Printf("Create: '%s'\n", name)
	return nil, fuse.ENOSYS
}

func (me *DropboxFs) OpenDir(name string, context *fuse.Context) (c []fuse.DirEntry, code fuse.Status) {
	path := get_path(name)
	log.Printf("OpenDir: '%s'\n", name)

	data, err := Cache.Metadata.Get(path)
	entry := fuse.DirEntry{}

	if data.IsDir && err == nil {
		for _, path := range data.Contents {
			entry.Name = NameFromPath(path)
			c = append(c, entry)
		}
		return c, fuse.OK
	}
	return nil, fuse.ENOENT
}

func (me *DropboxFs) Symlink(value string, linkName string, context *fuse.Context) (code fuse.Status) {
	log.Printf("Symlink: '%s'\n", linkName)
	return fuse.ENOSYS
}

func (me *DropboxFs) Readlink(name string, context *fuse.Context) (string, fuse.Status) {
	log.Printf("Readlink: '%s'\n", name)
	return "", fuse.ENOSYS
}
