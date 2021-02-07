package bili

import (
	"encoding/json"
)

//
//Get video information and description
//

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
	responseBytes, err := HttpGetWithParams(c.client, c.Endpoints.VideoViewUrl, request)
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
	responseBytes, err := HttpGetWithParams(c.client, c.Endpoints.VideoDescUrl, request)
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

//
//Like/Unlike a video
//

type LikeUnlikeOpCode int

const (
	VideoLike   LikeUnlikeOpCode = 1
	VideoUnlike LikeUnlikeOpCode = 2
)

type VideoLikeRequestAid struct {
	Aid  int              `url:"aid"`
	Like LikeUnlikeOpCode `url:"like"`
	Csrf string           `url:"csrf"`
}

type VideoLikeRequestBvid struct {
	Bvid string           `url:"bvid"`
	Like LikeUnlikeOpCode `url:"like"`
	Csrf string           `url:"csrf"`
}

type VideoLikeResponseCode int

const (
	VideoLikeSuccess        VideoLikeResponseCode = 0
	VideoLikeNotLoggedIn    VideoLikeResponseCode = -101
	VideoLikeWrongCSRF      VideoLikeResponseCode = -111
	VideoLikeBadRequest     VideoLikeResponseCode = -400
	VideoLikeVideoNotExists VideoLikeResponseCode = 10003
	VideoLikeCannotUnlike   VideoLikeResponseCode = 65004
	VideoLikeDuplicateLikes VideoLikeResponseCode = 65006
)

type VideoLikeResponse struct {
	Code    VideoLikeResponseCode `json:"code"`
	Message string                `json:"message"`
	Ttl     int                   `json:"ttl"`
}

func (c *Client) likeVideo(request interface{}) (VideoLikeResponse, error) {
	responseBody, _, err := HttpPostWithParams(c.client, c.Config.Endpoints.VideoLikeUrl, request)
	if err != nil {
		return VideoLikeResponse{}, err
	}
	response := VideoLikeResponse{}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return VideoLikeResponse{}, err
	}
	return response, nil
}

func (c *Client) LikeVideoByAid(aid int, like LikeUnlikeOpCode) (VideoLikeResponse, error) {
	csrf, err := c.getCookieValueByName("bili_jct")
	if err != nil {
		return VideoLikeResponse{}, err
	}
	return c.likeVideo(VideoLikeRequestAid{
		Aid:  aid,
		Like: like,
		Csrf: csrf,
	})
}

func (c *Client) LikeVideoByBvid(bvid string, like LikeUnlikeOpCode) (VideoLikeResponse, error) {
	csrf, err := c.getCookieValueByName("bili_jct")
	if err != nil {
		return VideoLikeResponse{}, err
	}
	return c.likeVideo(VideoLikeRequestBvid{
		Bvid: bvid,
		Like: like,
		Csrf: csrf,
	})
}

//
//Check a video has like or not
//

type LikeUnlikeStatusCode int

const (
	UnLiked LikeUnlikeStatusCode = 0
	Liked   LikeUnlikeStatusCode = 1
)

type VideoHasLikeResponse struct {
	Code    VideoLikeResponseCode `json:"code"`
	Message string                `json:"message"`
	Ttl     int                   `json:"ttl"`
	Data    LikeUnlikeStatusCode  `json:"data"`
}

func (c *Client) checkVideoLike(request interface{}) (VideoHasLikeResponse, error) {
	responseBody, err := HttpGetWithParams(c.client, c.Config.Endpoints.VideoCheckLikeUrl, request)
	if err != nil {
		return VideoHasLikeResponse{}, err
	}
	response := VideoHasLikeResponse{}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return VideoHasLikeResponse{}, err
	}
	return response, nil
}

func (c *Client) CheckVideoLikeByAid(aid int) (VideoHasLikeResponse, error) {
	return c.checkVideoLike(VideoRequestAid{Aid: aid})
}

func (c *Client) CheckVideoLikeByBvid(bvid string) (VideoHasLikeResponse, error) {
	return c.checkVideoLike(VideoRequestBvid{Bvid: bvid})
}

//
//Add coin to the video
//

type VideoAddCoinRequestAid struct {
	Aid        int    `url:"aid"`
	Multiply   int    `url:"multiply"`
	SelectLike int    `url:"select_like"`
	Csrf       string `url:"csrf"`
}

type VideoAddCoinRequestBvid struct {
	Bvid       string `url:"bvid"`
	Multiply   int    `url:"multiply"`
	SelectLike int    `url:"select_like"`
	Csrf       string `url:"csrf"`
}

type VideoAddCoinResponseCode int

const (
	VideoAddCoinSuccess           VideoAddCoinResponseCode = 0
	VideoAddCoinNotLoggedIn       VideoAddCoinResponseCode = -101
	VideoAddCoinAccountBanned     VideoAddCoinResponseCode = -102
	VideoAddCoinInsufficientCoin  VideoAddCoinResponseCode = -104
	VideoAddCoinWrongCSRF         VideoAddCoinResponseCode = -111
	VideoAddCoinBadRequest        VideoAddCoinResponseCode = -400
	VideoAddCoinVideoNotExists    VideoAddCoinResponseCode = 10003
	VideoAddCoinCannotToYourself  VideoAddCoinResponseCode = 34002
	VideoAddCoinInvalidMultiplier VideoAddCoinResponseCode = 34003
	VideoAddCoinIntervalToShort   VideoAddCoinResponseCode = 34004
	VideoAddCoinBeyondLimit       VideoAddCoinResponseCode = 34005
)

type VideoAddCoinResponse struct {
	Code    VideoAddCoinResponseCode `json:"code"`
	Message string                   `json:"message"`
	Ttl     int                      `json:"ttl"`
	Data    interface{}              `json:"data"`
}

func (c *Client) addCoinToVideo(request interface{}) (VideoAddCoinResponse, error) {
	responseBody, _, err := HttpPostWithParams(c.client, c.Config.Endpoints.VideoAddCoinUrl, request)
	if err != nil {
		return VideoAddCoinResponse{}, err
	}
	response := VideoAddCoinResponse{}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return VideoAddCoinResponse{}, err
	}
	return response, nil
}

func (c *Client) AddCoinToVideoByAid(aid int, multiply int, likeAsWell bool) (VideoAddCoinResponse, error) {
	selectLike := 0
	if likeAsWell {
		selectLike = 1
	}
	csrf, err := c.getCookieValueByName("bili_jct")
	if err != nil {
		return VideoAddCoinResponse{}, err
	}
	return c.addCoinToVideo(VideoAddCoinRequestAid{
		Aid:        aid,
		Multiply:   multiply,
		SelectLike: selectLike,
		Csrf:       csrf,
	})
}

func (c *Client) AddCoinToVideoByBvid(bvid string, multiply int, likeAsWell bool) (VideoAddCoinResponse, error) {
	selectLike := 0
	if likeAsWell {
		selectLike = 1
	}
	csrf, err := c.getCookieValueByName("bili_jct")
	if err != nil {
		return VideoAddCoinResponse{}, err
	}
	return c.addCoinToVideo(VideoAddCoinRequestBvid{
		Bvid:       bvid,
		Multiply:   multiply,
		SelectLike: selectLike,
		Csrf:       csrf,
	})
}
