package test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/scottferg/Dropbox-Go/dropbox"
	"github.com/sgodbold/dropbox-mnt/fs"
	"github.com/stretchr/testify/assert"
)

var upload_path = string("/testing/hello_world.txt")
var uri = dropbox.Uri{dropbox.RootDropbox, upload_path}

func setup() {
	fs.LoadConfig()
	fs.LoadSession()

	// Upload test file
	file, err := ioutil.ReadFile("./test_files/hello_world.txt")
	if err != nil {
		log.Fatalf("Setup Error: %v\n", err)
	}
	_, err = dropbox.UploadFile(fs.Session, file, uri, nil)
	if err != nil {
		log.Fatalf("Upload Error: %v\n", err)
	}
}

func teardown() {
	dropbox.Delete(fs.Session, uri, nil)
}

func TestMain(m *testing.M) {
	setup()
	status := m.Run()
	teardown()
	os.Exit(status)
}

func TestCacheInit(t *testing.T) {
	fs.CacheInit()
	metadata := fs.Metadata{Path: "/hello/world"}
	fs.Cache.Metadata.Data["hello"] = metadata
	assert.Exactly(t, metadata, fs.Cache.Metadata.Data["hello"])
}

func TestMetadataGet(t *testing.T) {
	assert := assert.New(t)
	fs.CacheInit()

	// Check that cache is currently empty
	assert.Exactly(fs.Metadata{}, fs.Cache.Metadata.Data[upload_path])

	metadata, err := fs.Cache.Metadata.Get(upload_path)

	assert.NoError(err)
	assert.NotEmpty(metadata)
	assert.NotEmpty(fs.Cache.Metadata.Data[upload_path])
}
