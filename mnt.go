package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/scottferg/Dropbox-Go/dropbox"
	"github.com/sgodbold/dropbox-mnt/fs"
)

// LoadConfig loads 'filename' into a global Configuration struct.
func loadConfig() (err error) {
	fs.Config = fs.Configuration{}

	file, open_err := os.Open(fs.CONFIG_NAME)
	if open_err != nil {
		file, _ = os.Create(fs.CONFIG_NAME)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&fs.Config)
	if open_err != nil || err != nil {
		err = initConfig(file)
	}
	return err
}

func initConfig(file *os.File) (err error) {
	// Get key and secret info from user
	fmt.Println("Enter the Following...")
	fmt.Printf("App Key: ")
	fmt.Scanln(&fs.Config.AppKey)
	fmt.Printf("App Secret: ")
	fmt.Scanln(&fs.Config.AppSecret)
	fmt.Printf("Access Type: ")
	fmt.Scanln(&fs.Config.AccessType)

	// Create dropbox session
	s := dropbox.Session{
		AppKey:     fs.Config.AppKey,
		AppSecret:  fs.Config.AppSecret,
		AccessType: fs.Config.AccessType,
	}
	s.ObtainRequestToken()

	// Obtain authorization URL and get token info from user
	url := dropbox.GenerateAuthorizeUrl(s.Token.Key, nil)
	fmt.Printf("Please visit this url and grant access: %s\n", url)
	fmt.Printf("Press 'Enter' when done...")
	var tmp string
	fmt.Scanln(&tmp)
	atoken, _ := s.ObtainAccessToken()
	fs.Config.TokenKey = atoken.Key
	fs.Config.TokenSecret = atoken.Secret

	// Write conf back to the config file.
	config_json, err := json.MarshalIndent(&fs.Config, "", "    ")
	if err != nil {
		return err
	}
	_, err = file.Write(config_json)

	return err
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("Usage:\n dropboxfs MOUNTPOINT")
	}
	err := loadConfig()
	if err != nil {
		log.Fatalf("Config fail: %v\n", err)
	}
	fs.CacheInit()
	fs.MountFs(flag.Arg(0))
}
