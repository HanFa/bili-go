package bili

import (
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
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
	fullEndpoint := endpoint + "?" + v.Encode()
	return HttpGet(fullEndpoint)
}

// https://gist.github.com/tonyhb/5819315
func structToMap(i interface{}) (values url.Values) {
	values = url.Values{}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		// You ca use tags here...
		// tag := typ.Field(i).Tag.Get("tagname")
		// Convert each type into a string for the url.Values string map
		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		}
		values.Set(typ.Field(i).Name, v)
	}
	return
}

func HttpPostWithParams(endpoint string, form interface{}) ([]byte, error) {
	v, err := query.Values(form)
	if err != nil {
		return nil, err
	}
	response, err := http.PostForm(endpoint, v)
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
