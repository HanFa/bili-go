package bili

import (
	"errors"
	"github.com/juju/persistent-cookiejar"
	"net/http"
)

type Client struct {
	Auth
	Zone
	config *Config
}

func setup(client *Client) (*Client, error) {
	// try loading cookies
	var err error
	var jar *cookiejar.Jar
	var jarPath = client.config.Cookies
	if jar, err = cookiejar.New(&cookiejar.Options{
		PublicSuffixList: nil,
		Filename:         jarPath,
	}); err != nil {
		// if cannot loading cookie file from
		// then create a new cookie jar
		if jar, err = cookiejar.New(nil); err != nil {
			return nil, err
		}
	}

	client.client.Jar = jar // set the cookie jar
	return client, nil
}

//New use the default config if the path to a config.json is not specified
func New() (client *Client, err error) {
	client = &Client{
		Auth: Auth{client: &http.Client{}},
	}

	client.config = &DefaultConfig
	return setup(client)
}

//New create a bili client if the path to a config.json file is specified
func NewWithConfig(configPath string) (client *Client, err error) {

	client = &Client{
		Auth: Auth{client: &http.Client{}},
	}

	var config *Config
	if config, err = LoadFromJSON(configPath); err != nil {
		return nil, errors.New("cannot load config: " + err.Error())
	}
	client.config = config
	return setup(client)
}
