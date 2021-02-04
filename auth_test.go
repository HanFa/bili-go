package bili

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestClient_GetCaptcha(t *testing.T) {
	response, err := client.GetCaptcha()
	assert.Nil(t, err, "cannot get captcha")
	assert.Equal(t, response.Code, CaptchaSuccess, "response code is not 0")
	assert.Equal(t, response.Data.Result.Success, 1, "Success is not 1")
}

func TestClient_GetPasswordSaltAndRSA(t *testing.T) {
	response, err := client.GetPasswordSaltAndRSA()
	assert.Nil(t, err, "cannot get password salt and rsa")
	assert.NotEmpty(t, response.Hash)
	assert.True(t, strings.HasPrefix(response.Key, "-----BEGIN PUBLIC KEY-----\n"))
}

func TestClient_EncryptPasswordWithSaltAndRSA(t *testing.T) {
	password, err := client.EncryptPasswordWithSaltAndRSA(
		"BiShi22332323", "8e0db05c46f4052c", "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDjb4V7EidX/ym28t2ybo0U6t0n\n6p4ej8VjqKHg100va6jkNbNTrLQqMCQCAYtXMXXp2Fwkk6WR+12N9zknLjf+C9sx\n/+l48mjUU8RqahiFD1XT/u2e0m2EN029OhCgkHx3Fc/KlFSIbak93EH/XlYis0w+\nXl69GV6klzgxW6d2xQIDAQAB\n-----END PUBLIC KEY-----")
	assert.Nil(t, err, "cannot encrypt password with salt and rsa")
	assert.Equal(t, password, "YgpjxAQ22pKa9socHIKPCZX0a/NS6Ng9Zzy+rp16b0LJGT6RHw2ERs3+ijCpG96PKTY1Baavwf0xgotmNvpl25l1KO5y4AjcqeWTzNTSVn6ejonBXGmBMybHHYawJ0aMPn1eDGpKrbI91mrF+h2x+fsnnpuZ1gheiYGzFmtshUc=")
}
