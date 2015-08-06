package fs

import (
	"github.com/scottferg/Dropbox-Go/dropbox"
)

func GetFile(path string) (file []byte, err error) {
	uri := MakeURI(path)
	params := dropbox.Parameters{}
	res, _, err := dropbox.GetFile(Session, uri, &params)
	if err != nil {
		return nil, err
	}
	return res, err
}

func GetMetadata(path string) (data Metadata, err error) {
	uri := MakeURI(path)
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
