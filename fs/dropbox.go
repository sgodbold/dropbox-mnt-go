package fs

import (
	"strings"

	"github.com/scottferg/Dropbox-Go/dropbox"
)

var Session dropbox.Session

// LoadSession starts a dropbox session and saves into Session
func LoadSession() (s dropbox.Session) {
	s = dropbox.Session{
		AppKey:     Config.AppKey,
		AppSecret:  Config.AppSecret,
		AccessType: Config.AccessType,
		Token: dropbox.AccessToken{
			Secret: Config.TokenSecret,
			Key:    Config.TokenKey,
		},
	}
	return s
}

// GetFile returns the contents of 'name' at 'path'
func GetFile(path string, name string) {
	return
}

// _nameFromPath returns the file or directory name at the end of 'path'
func _nameFromPath(path string) (name string) {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

// FilenamesInDir returns the filenames in 'path'
func FilenamesInDir(path string) (filenames []string, err error) {
	uri := dropbox.Uri{
		Root: dropbox.RootDropbox,
		Path: path,
	}
	params := dropbox.Parameters{List: "True"}

	res, err := dropbox.GetMetadata(Session, uri, &params)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(res.Contents); i++ {
		filenames = append(filenames, _nameFromPath(res.Contents[i].Path))
	}

	return filenames, err
}
