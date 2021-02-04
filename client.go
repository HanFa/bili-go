package bili

import "errors"

type Client struct {
	Config
}

func New(config string) (client *Client, err error) {
	client = &Client{}
	if err := client.LoadFromJSON(config); err != nil {
		return nil, errors.New("cannot load config: " + err.Error())
	}
	return client, nil
}
