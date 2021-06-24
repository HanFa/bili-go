package bili

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	cookiejar "github.com/juju/persistent-cookiejar"
)

type CaptchaResponseCode int

const (
	CaptchaSuccess CaptchaResponseCode = 0
)

type CaptchaGetResponse struct {
	Code CaptchaResponseCode `json:"code"`
	Data struct {
		Type   int `json:"type"`
		Result struct {
			Success   int    `json:"success"`
			Gt        string `json:"gt"`
			Challenge string `json:"challenge"`
			Key       string `json:"key"`
		} `json:"result"`
	} `json:"data"`
}

type SaltAndRsaResponse struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

type LoginRequest struct {
	CaptchaType int    `url:"captchaType"`
	Username    string `url:"username"`
	Password    string `url:"password"`
	Keep        bool   `url:"keep"`
	Key         string `url:"key"`
	Challenge   string `url:"challenge"`
	Validate    string `url:"validate"`
	Seccode     string `url:"seccode"`
}

type LoginResponseCode int

const (
	LoginSuccess                 LoginResponseCode = 0
	LoginBadRequest              LoginResponseCode = -400
	LoginWrongUsernameOrPassword LoginResponseCode = -629
	LoginEmptyUsernameOrPassword LoginResponseCode = -653
	LoginTimeout                 LoginResponseCode = -662
	LoginFieldMissing            LoginResponseCode = -2001
	LoginPhoneNumOrEmailNeeded   LoginResponseCode = -2100
	LoginEmailVerifiedNeed       LoginResponseCode = -2110
	LoginWrongPrivateKey         LoginResponseCode = 2400
	LoginWrongGeeTest            LoginResponseCode = 2406
)

type LoginResponse struct {
	Code    LoginResponseCode `json:"code"`
	Ts      int               `json:"ts"`
	Message string            `json:"message"`
	Data    interface{}       `json:"data"`
}

type Auth struct {
	Gt        string
	Challenge string
	Key       string

	Validate string
	Seccode  string

	Hash string
	RSA  string

	client *http.Client
}

// GetCaptcha retrieves the captcha content from Bilibili for DoCaptcha
func (c *Client) GetCaptcha() (CaptchaGetResponse, error) {
	responseBytes, err := HttpGet(c.client, c.config.Endpoints.CaptchaGetUrl)
	if err != nil {
		return CaptchaGetResponse{}, err
	}
	response := CaptchaGetResponse{}
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return CaptchaGetResponse{}, err
	}
	return response, nil
}

// DoCaptcha verifies the human interaction provided by https://kuresaru.github.io/geetest-validator/
func (c *Client) DoCaptcha() error {
	response, err := c.GetCaptcha()
	if err != nil {
		return errors.New("Cannot fetch challenge: " + err.Error())
	}

	c.Gt, c.Challenge, c.Key = response.Data.Result.Gt, response.Data.Result.Challenge, response.Data.Result.Key
	log.Printf("Gt: %s\nChallenge: %s\nKey: %s\n", c.Gt, c.Challenge, c.Key)
	log.Println("Please finish the test at https://kuresaru.github.io/geetest-validator/ and tell me validate and seccode")

	reader := bufio.NewReader(os.Stdin)
	log.Print("validate: ")
	validate, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	log.Print("seccode: ")
	seccode, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	validate = strings.Replace(validate, "\n", "", -1)
	seccode = strings.Replace(seccode, "\n", "", -1)

	c.Validate, c.Seccode = validate, seccode
	return nil
}

func (c *Client) GetPasswordSaltAndRSA() (SaltAndRsaResponse, error) {
	responseBytes, err := HttpGet(c.client, c.config.Endpoints.SaltAndRsaGetUrl)
	if err != nil {
		return SaltAndRsaResponse{}, err
	}
	response := SaltAndRsaResponse{}
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return SaltAndRsaResponse{}, err
	}
	return response, nil
}

func (c *Client) EncryptPasswordWithSaltAndRSA(plain, hash, key string) (string, error) {
	block, _ := pem.Decode([]byte(key))
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	pub := pubInterface.(*rsa.PublicKey)
	password, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(hash+plain))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(password), err
}

func (c *Client) IsSessionValid() (valid bool) {
	u, _ := url.Parse("http://bilibili.com")
	cookies := c.client.Jar.Cookies(u) // select all unexpired cookies
	if len(cookies) < 5 {
		return false
	}
	return true
}

func (c *Client) ClearSession() error {

	cookiesPath := ""
	if c.config.Cookies == "" {
		cookiesPath = cookiejar.DefaultCookieFile()
	} else {
		cookiesPath = c.config.Cookies
	}

	if err := os.Remove(cookiesPath); err != nil {
		return err
	}

	c.client.Jar, _ = cookiejar.New(nil)
	if err := c.client.Jar.(*cookiejar.Jar).Save(); err != nil {
		return err
	}
	return nil
}

// Login logins the user if the session is invalid or expired using the DoLogin
// provided with the username and plain password
func (c *Client) Login(username, plain string) (LoginResponse, error) {
	if !c.IsSessionValid() {
		return c.DoLogin(username, plain)
	}
	return LoginResponse{
		Code:    LoginSuccess,
		Ts:      0,
		Message: "successfully used cookies to authenticate",
		Data:    nil,
	}, nil
}

// DoLogin will interactively prompt the captcha verification by calling DoCaptcha,
// after that it logins with the provided with the username and plain password.
func (c *Client) DoLogin(username, plain string) (LoginResponse, error) {
	err := c.DoCaptcha()
	if err != nil {
		return LoginResponse{}, errors.New("Error when doing captcha: " + err.Error())
	}
	response, err := c.GetPasswordSaltAndRSA()
	if err != nil {
		return LoginResponse{}, errors.New("Error when fetching salt and RSA pubkey:" + err.Error())
	}
	c.Hash, c.RSA = response.Hash, response.Key
	password, err := c.EncryptPasswordWithSaltAndRSA(plain, c.Hash, c.RSA)
	if err != nil {
		return LoginResponse{}, errors.New("Error when encrypting password with salt and rsa: " + err.Error())
	}

	request := LoginRequest{
		CaptchaType: 6,
		Username:    username,
		Password:    password,
		Keep:        true,
		Key:         c.Key,
		Challenge:   c.Challenge,
		Validate:    c.Validate,
		Seccode:     c.Seccode,
	}

	responseBytes, _, err := HttpPostWithParams(c.client, c.config.Endpoints.LoginPostUrl, request)
	if err != nil {
		return LoginResponse{}, err
	}

	loginResponse := LoginResponse{}
	if err = json.Unmarshal(responseBytes, &loginResponse); err != nil {
		return LoginResponse{}, err
	}

	return loginResponse, nil
}
