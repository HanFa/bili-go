package bili

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_GetStreamUrlAvid_Flv(t *testing.T) {
	resp, err := client.GetStreamUrlAvid(99999999, 1, Stream720P60, StreamFlv, true) // 2P of this video
	assert.Nil(t, err, "Error occurs when fetching streaming url")
	assert.Greater(t, len(resp.Data.Durl), 0, "Urls for each partition should not be empty")
	assert.Equal(t, resp.Data.Durl[0].Size, 12259938)
}

func TestClient_GetStreamUrlBvid_Flv(t *testing.T) {
	resp, err := client.GetStreamUrlBvid("BV1y7411Q7Eq", 1, Stream720P60, StreamFlv, true) // 2P of this video
	assert.Nil(t, err, "Error occurs when fetching streaming url")
	assert.Greater(t, len(resp.Data.Durl), 0, "Urls for each partition should not be empty")
	assert.Equal(t, resp.Data.Durl[0].Size, 12259938)
}

func TestClient_GetStreamUrlBvid_LowRes(t *testing.T) {
	resp, err := client.GetStreamUrlBvid("BV1y7411Q7Eq", 1, Stream720P60, StreamLowResMp4, true) // 2P of this video
	assert.Nil(t, err, "Error occurs when fetching streaming url")
	assert.Greater(t, len(resp.Data.Durl), 0, "Urls for each partition should not be empty")
	assert.Equal(t, resp.Data.Durl[0].Size, 7381550)
}

func TestClient_GetStreamUrlAvid_Dash(t *testing.T) {
	resp, err := client.GetStreamUrlAvid(99999999, 1, Stream720P60, StreamDash, true) // 2P of this video
	assert.Nil(t, err, "Error occurs when fetching streaming url")
	assert.Greater(t, len(resp.Data.Dash.Video), 0, "No video stream")
	assert.Greater(t, len(resp.Data.Dash.Audio), 0, "No audio stream")
}
