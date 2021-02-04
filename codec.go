package bili

import (
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
)

func HttpGetWithParams(endpoint string, params interface{}) ([]byte, error) {
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	url := endpoint + "?" + v.Encode()
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, err
}
