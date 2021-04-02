package bili

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ConfigEndpoints struct {
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
	VideoChangeFavUrl    string `json:"video_change_fav_url"`
	StreamGetUrl         string `json:"stream_get_url"`
}

type Config struct {
	Cookies   string          `json:"cookies"`
	Endpoints ConfigEndpoints `json:"endpoints"`
}

var DefaultConfig Config = Config{
	Cookies: "./.bili-cookies",
	Endpoints: ConfigEndpoints{BaseUrl: "http://bilibili.com",
		CaptchaGetUrl:        "http://passport.bilibili.com/web/captcha/combine?plat=6",
		SaltAndRsaGetUrl:     "http://passport.bilibili.com/login?act=getkey",
		LoginPostUrl:         "http://passport.bilibili.com/web/login/v2",
		SplashBrandListUrl:   "http://app.bilibili.com/x/v2/splash/brand/list",
		VideoViewUrl:         "http://api.bilibili.com/x/web-interface/view",
		VideoDescUrl:         "http://api.bilibili.com/x/web-interface/archive/desc",
		VideoLikeUrl:         "http://api.bilibili.com/x/web-interface/archive/like",
		VideoCheckLikeUrl:    "http://api.bilibili.com/x/web-interface/archive/has/like",
		VideoAddCoinUrl:      "http://api.bilibili.com/x/web-interface/coin/add",
		VideoCheckHasCoinUrl: "http://api.bilibili.com/x/web-interface/archive/coins",
		VideoChangeFavUrl:    "http://api.bilibili.com/x/v3/fav/resource/deal",
		StreamGetUrl:         "http://api.bilibili.com/x/player/playurl"},
}

func LoadFromJSON(path string) (*Config, error) {
	config := new(Config)
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(byteValue, config); err != nil {
		return nil, err
	}
	return config, nil
}
