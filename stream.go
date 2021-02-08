package bili

import (
	"encoding/json"
	"errors"
	"fmt"
)

type StreamMode int

const (
	StreamFlv       = 0
	StreamLowResMp4 = 1
	StreamDash      = 16
)

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

type StreamUrlRequest struct {
	Cid        int                  `url:"cid"`
	Resolution StreamResolutionMode `url:"qn"`
	Mode       StreamMode           `url:"fnval"`
	Allow4K    int                  `url:"fourk"`
}

type StreamUrlRequestAid struct {
	Avid int `url:"avid"`
	StreamUrlRequest
}

type StreamUrlRequestBvid struct {
	Bvid string `url:"bvid"`
	StreamUrlRequest
}

type StreamUrlResponseCode int

const (
	StreamUrlSuccess        = 0
	StreamUrlBadRequest     = -400
	StreamUrlVideoNotExists = -404
)

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

func (c *Client) getStreamUrl(request interface{}) (StreamUrlResponse, error) {
	responseBytes, err := HttpGetWithParams(c.client, c.Endpoints.StreamGetUrl, request)
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
	return c.getStreamUrl(StreamUrlRequestAid{
		Avid: aid,
		StreamUrlRequest: StreamUrlRequest{
			Cid:        cid,
			Resolution: resolution,
			Mode:       mode,
			Allow4K:    boolToInt(allow4K),
		},
	})
}

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
	return c.getStreamUrl(StreamUrlRequestBvid{
		Bvid: bvid,
		StreamUrlRequest: StreamUrlRequest{
			Cid:        cid,
			Resolution: resolution,
			Mode:       mode,
			Allow4K:    boolToInt(allow4K),
		},
	})
}
