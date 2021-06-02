package bili

import (
	"errors"
	"fmt"
	cookiejar "github.com/juju/persistent-cookiejar"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/cheggaaa/pb/v3"
	"github.com/google/go-querystring/query"
)

type progress struct {
	enabled bool
	bar     *pb.ProgressBar
}

func (p *progress) Write(ch []byte) (n int, err error) {
	n = len(ch)
	p.bar.Add(n)
	return n, nil
}

func HttpGet(client *http.Client, endpoint string) ([]byte, error) {
	response, err := client.Get(endpoint)
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

func HttpGetAsFile(client *http.Client, endpoint string, path string, showProgress bool, progressWriter io.Writer) error {
	tmpPath := fmt.Sprintf("%s.tmp", path)
	out, err := os.Create(tmpPath)
	if err != nil {
		return err
	}
	response, err := client.Get(endpoint)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return errors.New("get response does not return 200")
	}
	progressBar := progress{showProgress, pb.StartNew(int(response.ContentLength))}
	teeReader := io.TeeReader(response.Body, &progressBar)
	if progressWriter != nil {
		teeReader = io.TeeReader(teeReader, progressWriter)
	}
	_, err = io.Copy(out, teeReader)
	progressBar.bar.Finish()
	if err != nil {
		return err
	}
	if err = os.Rename(tmpPath, path); err != nil {
		return err
	}
	return nil
}

func HttpGetWithParams(client *http.Client, endpoint string, params interface{}) ([]byte, error) {
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	fullEndpoint := endpoint + "?" + v.Encode()
	return HttpGet(client, fullEndpoint)
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

func HttpPostWithParams(client *http.Client, endpoint string, form interface{}) ([]byte, []*http.Cookie, error) {
	return HttpPostWithParamsReferer(client, endpoint, form, "")
}

func HttpPostWithParamsReferer(client *http.Client, endpoint string, form interface{}, referer string) ([]byte, []*http.Cookie, error) {
	v, err := query.Values(form)
	if err != nil {
		return nil, nil, err
	}
	request, err := http.NewRequest("POST", endpoint, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, nil, err
	}
	if len(referer) > 0 {
		request.Header.Set("Referer", referer)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}
	u, _ := url.Parse("http://bilibili.com")
	client.Jar.SetCookies(u, response.Cookies())
	if err := client.Jar.(*cookiejar.Jar).Save(); err != nil {
		log.Printf("HttpPostWithParamsReferer: cannot save cookie %v\n", err)
	}
	return responseBody, response.Cookies(), err
}
