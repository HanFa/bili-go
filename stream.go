package bili

import (
	"encoding/json"
	"errors"
	"fmt"
)

// StreamMode specifies the mode of streaming either flv, low-res mp4 or dash
type StreamMode int

// List of supported StreamMode
const (
	StreamFlv       = 0
	StreamLowResMp4 = 1
	StreamDash      = 16
)

// StreamResolutionMode specifies the resolution of the streaming
type StreamResolutionMode int

const (
	Stream240P      StreamResolutionMode = 6
	Stream360P      StreamResolutionMode = 16
	Stream480P      StreamResolutionMode = 32
	Stream720P      StreamResolutionMode = 64
	Stream720P60    StreamResolutionMode = 74
	Stream1080P     StreamResolutionMode = 80
	Stream1080PPlus StreamResolutionMode = 112
	Stream1080P60   StreamResolutionMode = 116
	Stream4K        StreamResolutionMode = 120
)

const (
	StreamAudio64K  StreamResolutionMode = 30216
	StreamAudio132K StreamResolutionMode = 30232
	StreamAudio192K StreamResolutionMode = 30280
)

// streamUrlRequest represents a request payload to retrieve stream url
type streamUrlRequest struct {
	Cid        int                  `url:"cid"`
	Resolution StreamResolutionMode `url:"qn"`
	Mode       StreamMode           `url:"fnval"`
	Allow4K    int                  `url:"fourk"`
}

// streamUrlRequestAid represents a streamUrlRequest specified by video Aid
type streamUrlRequestAid struct {
	Avid int `url:"avid"`
	streamUrlRequest
}

// streamUrlRequestBvid represents a streamUrlRequest specified by video Bvid
type streamUrlRequestBvid struct {
	Bvid string `url:"bvid"`
	streamUrlRequest
}

// StreamUrlResponseCode represents the status code of retrieving stream URL response
type StreamUrlResponseCode int

const (
	StreamUrlSuccess        = 0
	StreamUrlBadRequest     = -400
	StreamUrlVideoNotExists = -404
)

// StreamDashInfo represents the dash information of the retrieved stream
type StreamDashInfo struct {
	Id        StreamResolutionMode `json:"id"`
	BaseUrl   string               `json:"base_url"`
	BackupUrl []string             `json:"backup_url"`
	Bandwidth int                  `json:"bandwidth"`
	MimeType  string               `json:"mime_type"`
	Codecs    string               `json:"codecs"`
	Width     int                  `json:"width"`
	Height    int                  `json:"height"`
	FrameRate string               `json:"frame_rate"`
}

// StreamUrlResponse represents the response payload of the stream url retrieval
type StreamUrlResponse struct {
	Code    StreamUrlResponseCode `json:"code"`
	Message string                `json:"message"`
	Ttl     int                   `json:"ttl"`
	Data    struct {
		Quality    StreamResolutionMode `json:"quality"`
		Format     string               `json:"format"`
		TimeLength int                  `json:"timelength"`
		Durl       []struct {
			Order     int      `json:"order"`
			Length    int      `json:"length"`
			Size      int      `json:"size"`
			Url       string   `json:"url"`
			BackupUrl []string `json:"backup_url"`
		} `json:"durl"`
		Dash struct {
			Video []StreamDashInfo `json:"video"`
			Audio []StreamDashInfo `json:"audio"`
		} `json:"dash"`
	} `json:"data"`
}

// getStreamUrl retrieves the stream url with the request payload
func (c *Client) getStreamUrl(request interface{}) (StreamUrlResponse, error) {
	responseBytes, err := HttpGetWithParams(c.client, c.config.Endpoints.StreamGetUrl, request)
	if err != nil {
		return StreamUrlResponse{}, err
	}
	response := StreamUrlResponse{}
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return StreamUrlResponse{}, err
	}
	return response, nil
}

func boolToInt(b bool) int {
	res := 0
	if b {
		res = 1
	}
	return res
}

// GetStreamUrlAvid retrieves the stream url with the specified Avid
func (c *Client) GetStreamUrlAvid(aid int, page int, resolution StreamResolutionMode,
	mode StreamMode, allow4K bool) (StreamUrlResponse, error) {
	info, err := c.GetVideoInfoByAid(aid)
	if err != nil {
		return StreamUrlResponse{}, err
	}
	pages := info.Data.Pages
	if page < 0 || page >= len(pages) {
		return StreamUrlResponse{},
			errors.New(fmt.Sprintf("Invalid page num %d, this video has %d page", page, len(pages)))
	}
	cid := pages[page].Cid
	return c.getStreamUrl(streamUrlRequestAid{
		Avid: aid,
		streamUrlRequest: streamUrlRequest{
			Cid:        cid,
			Resolution: resolution,
			Mode:       mode,
			Allow4K:    boolToInt(allow4K),
		},
	})
}

// GetStreamUrlBvid retrieves the stream url with the specified Bvid
func (c *Client) GetStreamUrlBvid(bvid string, page int, resolution StreamResolutionMode,
	mode StreamMode, allow4K bool) (StreamUrlResponse, error) {
	info, err := c.GetVideoInfoByBvid(bvid)
	if err != nil {
		return StreamUrlResponse{}, err
	}
	pages := info.Data.Pages
	if page < 0 || page >= len(pages) {
		return StreamUrlResponse{},
			errors.New(fmt.Sprintf("Invalid page num %d, this video has %d page", page, len(pages)))
	}
	cid := pages[page].Cid
	return c.getStreamUrl(streamUrlRequestBvid{
		Bvid: bvid,
		streamUrlRequest: streamUrlRequest{
			Cid:        cid,
			Resolution: resolution,
			Mode:       mode,
			Allow4K:    boolToInt(allow4K),
		},
	})
}
