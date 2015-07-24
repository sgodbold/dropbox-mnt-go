package fs

import (
	"strings"

	"github.com/scottferg/Dropbox-Go/dropbox"
)

var Session dropbox.Session

func LoadSession() {
	Session = dropbox.Session{
		AppKey:     Config.AppKey,
		AppSecret:  Config.AppSecret,
		AccessType: Config.AccessType,
		Token: dropbox.AccessToken{
			Secret: Config.TokenSecret,
			Key:    Config.TokenKey,
		},
	}
	return
}

func NameFromPath(path string) (name string) {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

func DirInfo(path string) (info dropbox.Metadata, err error) {
	uri := dropbox.Uri{
		Root: dropbox.RootDropbox,
		Path: path,
	}
	params := dropbox.Parameters{List: "True"}

	res, err := dropbox.GetMetadata(Session, uri, &params)

	return res, err
}
