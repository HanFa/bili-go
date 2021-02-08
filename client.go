package bili

import (
	"errors"
	"net/http"
	"net/http/cookiejar"
)

type Client struct {
	Auth
	Config
	Zone
}

func New(config string) (client *Client, err error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client = &Client{
		Auth: Auth{client: &http.Client{Jar: jar}},
	}
	if err := client.LoadFromJSON(config); err != nil {
		return nil, errors.New("cannot load config: " + err.Error())
	}
	return client, nil
}
