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

// NOTE: might move in the future
func testMakeURI(t *testing.T) {
	test_uri := fs.MakeURI(upload_path)
	assert.Exactly(t, uri, test_uri)
}

// NOTE: might move in the future
func testNameFromPath(t *testing.T) {
	path := "/test/path/to/file"
	name := fs.NameFromPath(path)
	assert.Exactly(t, "file", name)
}

func testGetFile(t *testing.T) {
	local_file, err := ioutil.ReadFile("./test_files/hello_world.txt")
	file, err := fs.GetFile(upload_path)
	assert.NoError(t, err)
	assert.Exactly(t, local_file, file)
}

func testGetMetadata(t *testing.T) {
	assert := assert.New(t)
	metadata, err := fs.GetMetadata(upload_path)
	assert.NoError(err)
	assert.Exactly(upload_path, metadata.Path)
	assert.False(metadata.IsDir)
	assert.Empty(metadata.Contents)
}
