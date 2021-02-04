package bili

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"
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

type Auth struct {
	Gt        string
	Challenge string
	Key       string

	Validate string
	Seccode  string
}

func (c *Client) GetCaptcha() (CaptchaGetResponse, error) {
	responseBytes, err := HttpGet(c.Config.Endpoints.CaptchaGetUrl)
	if err != nil {
		return CaptchaGetResponse{}, nil
	}
	response := CaptchaGetResponse{}
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return CaptchaGetResponse{}, nil
	}
	return response, nil
}

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
