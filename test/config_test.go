package test

import (
	"testing"

	"github.com/sgodbold/dropbox-mnt/fs"
	"github.com/stretchr/testify/assert"
)

func TestConfigSuccess(t *testing.T) {
	assert := assert.New(t)
	err := fs.LoadConfig("./test_configs/good_config.json")
	assert.NoError(err)
	assert.NotEmpty(fs.Config.AppKey)
	assert.NotEmpty(fs.Config.AppSecret)
	assert.NotEmpty(fs.Config.AccessType)
	assert.NotEmpty(fs.Config.TokenSecret)
	assert.NotEmpty(fs.Config.TokenKey)
}

func TestConfigMissing(t *testing.T) {
	err := fs.LoadConfig("./test_configs/missing_config.json")
	assert.NoError(t, err)
	assert.Empty(t, fs.Config.AppSecret)
}

func TestConfigSyntax(t *testing.T) {
	err := fs.LoadConfig("./test_configs/bad_syntax_config.json")
	assert.Error(t, err)
}
