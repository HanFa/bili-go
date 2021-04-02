package bili

import (
	"net/url"
)

func (c *Client) getCookieValueByName(name string) (string, error) {
	u, err := url.Parse(c.config.Endpoints.BaseUrl)
	if err != nil {
		return "", err
	}
	cookies := c.Auth.client.Jar.Cookies(u)
	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie.Value, nil
		}
	}
	return "", nil
}
