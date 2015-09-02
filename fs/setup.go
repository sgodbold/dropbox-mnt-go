package fs

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
	"github.com/scottferg/Dropbox-Go/dropbox"
)

func CacheInit() {
	Cache.Metadata.Data = make(map[string]Metadata)
}

func InitConfig(file *os.File) (err error) {
	// Get key and secret info from user
	fmt.Println("Enter the Following...")
	fmt.Printf("App Key: ")
	fmt.Scanln(&Config.AppKey)
	fmt.Printf("App Secret: ")
	fmt.Scanln(&Config.AppSecret)
	fmt.Printf("Access Type: ")
	fmt.Scanln(&Config.AccessType)

	// Create dropbox session
	s := dropbox.Session{
		AppKey:     Config.AppKey,
		AppSecret:  Config.AppSecret,
		AccessType: Config.AccessType,
	}
	s.ObtainRequestToken()

	// Obtain authorization URL and get token info from user
	url := dropbox.GenerateAuthorizeUrl(s.Token.Key, nil)
	fmt.Printf("Please visit this url and grant access: %s\n", url)
	fmt.Printf("Press 'Enter' when done...")
	var tmp string
	fmt.Scanln(&tmp)
	atoken, _ := s.ObtainAccessToken()
	Config.TokenKey = atoken.Key
	Config.TokenSecret = atoken.Secret
	Config.DropBoxMnt = "/"

	// Write conf back to the config file.
	config_json, err := json.MarshalIndent(&Config, "", "    ")
	if err != nil {
		return err
	}
	_, err = file.Write(config_json)

	return err
}

func LoadConfig() (err error) {
	Config = Configuration{}

	file, open_err := os.Open(CONFIG_NAME)
	if open_err != nil {
		file, _ = os.Create(CONFIG_NAME)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if open_err != nil || err != nil {
		err = InitConfig(file)
	}
	return err
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

func MountFs(path string) {
	LoadSession()
	CacheInit()
	nfs := pathfs.NewPathNodeFs(&DropboxFs{FileSystem: pathfs.NewDefaultFileSystem()}, nil)
	server, _, err := nodefs.MountRoot(path, nfs.Root(), nil)
	if err != nil {
		log.Fatalf("Mount fail: %v\n", err)
	}
	server.Serve()
}
