package fs

import (
	"strings"

	"github.com/scottferg/Dropbox-Go/dropbox"
)

var Session dropbox.Session

type Metadata struct {
	Hash     string
	Bytes    int
	Modified string
	Path     string
	IsDir    bool
	Contents []string // list of paths
}

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

func GetMetadata(path string) (data Metadata, err error) {
	uri := dropbox.Uri{
		Root: dropbox.RootDropbox,
		Path: path,
	}
	params := dropbox.Parameters{List: "True"}
	res, err := dropbox.GetMetadata(Session, uri, &params)
	if err != nil {
		return Metadata{}, err
	}

	// Convert response to Metadata type
	data.Hash = res.Hash
	data.Bytes = res.Bytes
	data.Path = res.Path
	data.IsDir = res.IsDir
	for _, file := range res.Contents {
		data.Contents = append(data.Contents, file.Path)
	}
	return data, err
}
