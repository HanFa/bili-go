package bili

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var client *Client

func TestMain(m *testing.M) {
	client, _ = New()
	_ = client.ClearSession()
	m.Run()
}

func TestClient_NotNil(t *testing.T) {
	assert.NotNil(t, client, "client cannot be nil")
}
