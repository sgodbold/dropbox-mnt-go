package fs

type Configuration struct {
	AppKey      string
	AppSecret   string
	AccessType  string
	TokenSecret string
	TokenKey    string
}

var Config Configuration
