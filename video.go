package bili

import (
	"encoding/json"
)

type VideoRequestAid struct {
	Aid int `url:"aid"`
}

type VideoRequestBvid struct {
	Bvid string `url:"bvid"`
}

type VideoResponseCode int

const (
	VideoSuccess      VideoResponseCode = 0
	VideoBadRequest   VideoResponseCode = -400
	VideoNoPermission VideoResponseCode = -403
	VideoNotFound     VideoResponseCode = -404
	VideoNotVisible   VideoResponseCode = 62002
)

type VideoInfoResponse struct {
	Code    VideoResponseCode `json:"code"`
	Message string            `json:"message"`
	Ttl     int               `json:"ttl"`
	Data    struct {
		Bvid      string `json:"bvid"`
		Aid       int    `json:"aid"`
		Videos    int    `json:"videos"`
		Tid       int    `json:"tid"`
		Tname     string `json:"tname"`
		Copyright int    `json:"copyright"`
		Pic       string `json:"pic"`
		Title     string `json:"title"`
		Pubdate   int    `json:"pubdate"`
		Ctime     int    `json:"ctime"`
		Desc      string `json:"desc"`
		State     int    `json:"state"`
		Attribute int    `json:"attribute"`
		Duration  int    `json:"duration"`
	} `json:"data"`
}

type VideoDescResponse struct {
	Code    VideoResponseCode `json:"code"`
	Message string            `json:"message"`
	Ttl     int               `json:"ttl"`
	Data    string            `json:"data"`
}

func (c *Client) getVideoInfo(request interface{}) (VideoInfoResponse, error) {
	responseBytes, err := HttpGetWithParams(c.Endpoints.VideoViewUrl, request)
	if err != nil {
		return VideoInfoResponse{}, err
	}
	response := VideoInfoResponse{}
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return VideoInfoResponse{}, err
	}
	return response, nil
}

func (c *Client) GetVideoInfoByAid(aid int) (VideoInfoResponse, error) {
	return c.getVideoInfo(VideoRequestAid{Aid: aid})
}

func (c *Client) GetVideoInfoByBvid(bvid string) (VideoInfoResponse, error) {
	return c.getVideoInfo(VideoRequestBvid{Bvid: bvid})
}

func (c *Client) getVideoDescription(request interface{}) (VideoDescResponse, error) {
	responseBytes, err := HttpGetWithParams(c.Endpoints.VideoDescUrl, request)
	if err != nil {
		return VideoDescResponse{}, err
	}
	response := VideoDescResponse{}
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return VideoDescResponse{}, err
	}
	return response, nil
}

func (c *Client) GetVideoDescriptionByAid(aid int) (VideoDescResponse, error) {
	return c.getVideoDescription(VideoRequestAid{Aid: aid})
}

func (c *Client) GetVideoDescriptionByBvid(bvid string) (VideoDescResponse, error) {
	return c.getVideoDescription(VideoRequestBvid{Bvid: bvid})
}
