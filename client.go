package bili

import (
	"errors"
	"github.com/juju/persistent-cookiejar"
	"net/http"
)

type Client struct {
	Auth
	Config
	Zone
}

func New(config string) (client *Client, err error) {

	client = &Client{
		Auth: Auth{client: &http.Client{}},
	}
	if err := client.LoadFromJSON(config); err != nil {
		return nil, errors.New("cannot load config: " + err.Error())
	}

	// try loading cookies
	var jar *cookiejar.Jar
	var jarPath = client.Config.Cookies
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
