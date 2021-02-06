package bili

import (
	"errors"
	"net/http"
)

type Client struct {
	Auth
	Config
}

func New(config string) (client *Client, err error) {
	client = &Client{
		Auth: Auth{client: &http.Client{}},
	}
	if err := client.LoadFromJSON(config); err != nil {
		return nil, errors.New("cannot load config: " + err.Error())
	}
	return client, nil
}
