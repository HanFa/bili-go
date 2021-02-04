package bili

import (
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
)

func HttpGet(endpoint string) ([]byte, error) {
	response, err := http.Get(endpoint)
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

func HttpGetWithParams(endpoint string, params interface{}) ([]byte, error) {
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	url := endpoint + "?" + v.Encode()
	return HttpGet(url)
}
