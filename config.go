package bili

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Endpoints struct {
		BaseUrl              string `json:"base_url"`
		CaptchaGetUrl        string `json:"captcha_get_url"`
		SaltAndRsaGetUrl     string `json:"salt_and_rsa_get_url"`
		LoginPostUrl         string `json:"login_post_url"`
		SplashBrandListUrl   string `json:"splash_brand_list_url"`
		VideoViewUrl         string `json:"video_view_url"`
		VideoDescUrl         string `json:"video_desc_url"`
		VideoLikeUrl         string `json:"video_like_url"`
		VideoCheckLikeUrl    string `json:"video_check_like_url"`
		VideoAddCoinUrl      string `json:"video_add_coin_url"`
		VideoCheckHasCoinUrl string `json:"video_check_has_coin_url"`
	} `json:"endpoints"`
}

func (c *Config) LoadFromJSON(path string) (err error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(byteValue, c); err != nil {
		return err
	}
	return nil
}
