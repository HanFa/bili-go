package bili

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_GetCaptcha(t *testing.T) {
	response, err := client.GetCaptcha()
	assert.Nil(t, err, "cannot get captcha")
	assert.Equal(t, response.Code, CaptchaSuccess, "response code is not 0")
	assert.Equal(t, response.Data.Result.Success, 1, "Success is not 1")
}
