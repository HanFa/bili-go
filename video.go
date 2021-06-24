package bili

import (
	"encoding/json"
)

//
// Get video information and description
//

// VideoRequestAid represents the request payload for the video information specified by Aid
type VideoRequestAid struct {
	Aid int `url:"aid"`
}

// VideoRequestBvid represents the request payload for the video information specified by Bvid
type VideoRequestBvid struct {
	Bvid string `url:"bvid"`
}

// VideoResponseCode is the status code of the response for fetching the video information
type VideoResponseCode int

const (
	VideoSuccess      VideoResponseCode = 0
	VideoBadRequest   VideoResponseCode = -400
	VideoNoPermission VideoResponseCode = -403
	VideoNotFound     VideoResponseCode = -404
	VideoNotVisible   VideoResponseCode = 62002
)

// VideoState is the state of the retrieved video information
type VideoState int

// VideoDimension is the dimension for the 1P video
type VideoDimension struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Rotate int `json:"rotate"`
}

// VideoInfoResponse represents the retrieved video information
type VideoInfoResponse struct {
	Code    VideoResponseCode `json:"code"`
	Message string            `json:"message"`
	Ttl     int               `json:"ttl"`
	Data    struct {
		Bvid        string     `json:"bvid"`
		Aid         int        `json:"aid"`
		Videos      int        `json:"videos"`
		Tid         int        `json:"tid"`
		Tname       string     `json:"tname"`
		Copyright   int        `json:"copyright"`
		Pic         string     `json:"pic"`
		Title       string     `json:"title"`
		Pubdate     int        `json:"pubdate"`
		Ctime       int        `json:"ctime"`
		Desc        string     `json:"desc"`
		State       VideoState `json:"state"`
		Attribute   int        `json:"attribute"`
		Duration    int        `json:"duration"`
		Forward     int        `json:"forward"`
		MissionId   int        `json:"mission_id"`
		RedirectUrl string     `json:"redirect_url"`
		Rights      struct {
			Bp            int `json:"bp"`
			Elec          int `json:"elec"`
			Download      int `json:"download"`
			Movie         int `json:"movie"`
			Pay           int `json:"pay"`
			Hd5           int `json:"hd5"`
			NoReprint     int `json:"no_reprint"`
			Autoplay      int `json:"autoplay"`
			UcgPlay       int `json:"ucg_play"`
			IsCooperation int `json:"is_cooperation"`
			UgcPayPreview int `json:"ugc_pay_preview"`
			NoBackground  int `json:"no_background"`
		} `json:"rights"`
		Owner struct {
			Mid  int    `json:"mid"`
			Name string `json:"name"`
			Face string `json:"face"`
		} `json:"owner"`
		Stat struct {
			Aid        int    `json:"aid"`
			View       int    `json:"view"`
			Danmaku    int    `json:"danmaku"`
			Reply      int    `json:"reply"`
			Favorite   int    `json:"favorite"`
			Coin       int    `json:"coin"`
			Share      int    `json:"share"`
			NowRank    int    `json:"now_rank"`
			HisRank    int    `json:"his_rank"`
			Like       int    `json:"like"`
			Dislike    int    `json:"dislike"`
			Evaluation string `json:"evaluation"`
		} `json:"stat"`
		Dynamic   string         `json:"dynamic"`
		Cid       int            `json:"cid"`
		Dimension VideoDimension `json:"dimension"`
		NoCache   bool           `json:"no_cache"`
		Pages     []struct {
			Cid       int            `json:"cid"`
			Page      int            `json:"page"`
			From      string         `json:"from"`
			Part      string         `json:"part"`
			Duration  int            `json:"duration"`
			Vid       string         `json:"vid"`
			Weblink   string         `json:"weblink"`
			Dimension VideoDimension `json:"dimension"`
		} `json:"pages"`
		Subtitle struct {
			AllowSubmit bool `json:"allow_submit"`
			List        []struct {
				Id          int    `json:"id"`
				Lan         string `json:"lan"`
				LanDoc      string `json:"lan_doc"`
				IsLock      bool   `json:"is_lock"`
				AuthorMid   int    `json:"author_mid"`
				SubtitleUrl string `json:"subtitle_url"`
			} `json:"list"`
		} `json:"subtitle"`
		Staff []struct{} `json:"staff"`
	} `json:"data"`
}

// VideoDescResponse represents the response of the retrieved video description
type VideoDescResponse struct {
	Code    VideoResponseCode `json:"code"`
	Message string            `json:"message"`
	Ttl     int               `json:"ttl"`
	Data    string            `json:"data"`
}

// getVideoInfo retrieves the video information
func (c *Client) getVideoInfo(request interface{}) (VideoInfoResponse, error) {
	responseBytes, err := HttpGetWithParams(c.client, c.config.Endpoints.VideoViewUrl, request)
	if err != nil {
		return VideoInfoResponse{}, err
	}
	response := VideoInfoResponse{}
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return VideoInfoResponse{}, err
	}
	return response, nil
}

// GetVideoInfoByAid retrieves the video information specified by the Aid
func (c *Client) GetVideoInfoByAid(aid int) (VideoInfoResponse, error) {
	return c.getVideoInfo(VideoRequestAid{Aid: aid})
}

// GetVideoInfoByBvid retrieves the video information specified by the Bvid
func (c *Client) GetVideoInfoByBvid(bvid string) (VideoInfoResponse, error) {
	return c.getVideoInfo(VideoRequestBvid{Bvid: bvid})
}

// getVideoDescription retrieves the video description
func (c *Client) getVideoDescription(request interface{}) (VideoDescResponse, error) {
	responseBytes, err := HttpGetWithParams(c.client, c.config.Endpoints.VideoDescUrl, request)
	if err != nil {
		return VideoDescResponse{}, err
	}
	response := VideoDescResponse{}
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return VideoDescResponse{}, err
	}
	return response, nil
}

// GetVideoDescriptionByAid retrieves the video description specified by the Aid
func (c *Client) GetVideoDescriptionByAid(aid int) (VideoDescResponse, error) {
	return c.getVideoDescription(VideoRequestAid{Aid: aid})
}

// GetVideoDescriptionByBvid retrieves the video description specified by the Bvid
func (c *Client) GetVideoDescriptionByBvid(bvid string) (VideoDescResponse, error) {
	return c.getVideoDescription(VideoRequestBvid{Bvid: bvid})
}

//
// Like/Unlike a video
//

// LikeUnlikeOpCode is the operation code for Like/Unlike. 1 for like and 2 for unlike
type LikeUnlikeOpCode int

const (
	VideoLike   LikeUnlikeOpCode = 1
	VideoUnlike LikeUnlikeOpCode = 2
)

// VideoLikeRequestAid is the like/unlike request to video specified by the Aid
type VideoLikeRequestAid struct {
	Aid  int              `url:"aid"`
	Like LikeUnlikeOpCode `url:"like"`
	Csrf string           `url:"csrf"`
}

// VideoLikeRequestBvid is the like/unlike request to video specified by the Bvid
type VideoLikeRequestBvid struct {
	Bvid string           `url:"bvid"`
	Like LikeUnlikeOpCode `url:"like"`
	Csrf string           `url:"csrf"`
}

// VideoLikeResponseCode is the status code for the like/unlike response
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

// VideoLikeResponse is the response for like/unlike a video
type VideoLikeResponse struct {
	Code    VideoLikeResponseCode `json:"code"`
	Message string                `json:"message"`
	Ttl     int                   `json:"ttl"`
}

// likeVideo likes or unlikes a video
func (c *Client) likeVideo(request interface{}) (VideoLikeResponse, error) {
	responseBody, _, err := HttpPostWithParams(c.client, c.config.Endpoints.VideoLikeUrl, request)
	if err != nil {
		return VideoLikeResponse{}, err
	}
	response := VideoLikeResponse{}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return VideoLikeResponse{}, err
	}
	return response, nil
}

// LikeVideoByAid likes or unlikes a video specified by Aid
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

// LikeVideoByBvid likes or unlikes a video specified by Bvid
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
// Check a video has like or not
//

// LikeUnlikeStatusCode is the status code of a like/unlike operation
type LikeUnlikeStatusCode int

const (
	UnLiked LikeUnlikeStatusCode = 0
	Liked   LikeUnlikeStatusCode = 1
)

// VideoHasLikeResponse is the response of checking if the logged in user has liked a video or not
type VideoHasLikeResponse struct {
	Code    VideoLikeResponseCode `json:"code"`
	Message string                `json:"message"`
	Ttl     int                   `json:"ttl"`
	Data    LikeUnlikeStatusCode  `json:"data"`
}

// checkVideoLike checks if the logged in user has liked a video or not
func (c *Client) checkVideoLike(request interface{}) (VideoHasLikeResponse, error) {
	responseBody, err := HttpGetWithParams(c.client, c.config.Endpoints.VideoCheckLikeUrl, request)
	if err != nil {
		return VideoHasLikeResponse{}, err
	}
	response := VideoHasLikeResponse{}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return VideoHasLikeResponse{}, err
	}
	return response, nil
}

// CheckVideoLikeByAid checks if the logged in user has liked a video or not specified by Aid
func (c *Client) CheckVideoLikeByAid(aid int) (VideoHasLikeResponse, error) {
	return c.checkVideoLike(VideoRequestAid{Aid: aid})
}

// CheckVideoLikeByBvid checks if the logged in user has liked a video or not specified by Bvid
func (c *Client) CheckVideoLikeByBvid(bvid string) (VideoHasLikeResponse, error) {
	return c.checkVideoLike(VideoRequestBvid{Bvid: bvid})
}

//
// Add coin to the video
//

// VideoAddCoinRequestAid is a request payload to add coin to a video specified by Aid
type VideoAddCoinRequestAid struct {
	Aid        int    `url:"aid"`
	Multiply   int    `url:"multiply"`
	SelectLike int    `url:"select_like"`
	Csrf       string `url:"csrf"`
}

// VideoAddCoinRequestBvid is a request payload to add coin to a video specified by Bvid
type VideoAddCoinRequestBvid struct {
	Bvid       string `url:"bvid"`
	Multiply   int    `url:"multiply"`
	SelectLike int    `url:"select_like"`
	Csrf       string `url:"csrf"`
}

// VideoAddCoinResponseCode is the status code of the response for adding coin to a video
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

// VideoAddCoinResponse is the response for adding coin to a video
type VideoAddCoinResponse struct {
	Code    VideoAddCoinResponseCode `json:"code"`
	Message string                   `json:"message"`
	Ttl     int                      `json:"ttl"`
	Data    interface{}              `json:"data"`
}

// addCoinToVideo adds a coin to a video
func (c *Client) addCoinToVideo(request interface{}) (VideoAddCoinResponse, error) {
	responseBody, _, err := HttpPostWithParams(c.client, c.config.Endpoints.VideoAddCoinUrl, request)
	if err != nil {
		return VideoAddCoinResponse{}, err
	}
	response := VideoAddCoinResponse{}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return VideoAddCoinResponse{}, err
	}
	return response, nil
}

// AddCoinToVideoByAid adds a coin to a video specified by Aid
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

// AddCoinToVideoByBvid adds a coin to a video specified by Bvid
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

//
// Check a video has received coins from you or not
//

// VideoHasCoinsResponse is the response of checking if the logged-in user has sent coins to this video
type VideoHasCoinsResponse struct {
	Code    VideoAddCoinResponseCode `json:"code"`
	Message string                   `json:"message"`
	Ttl     int                      `json:"ttl"`
	Data    struct {
		Multiply int `json:"multiply"`
	} `json:"data"`
}

// checkVideoHasCoins checks if the logged-in user has sent coins to this video
func (c *Client) checkVideoHasCoins(request interface{}) (VideoHasCoinsResponse, error) {
	responseBody, err := HttpGetWithParams(c.client, c.config.Endpoints.VideoCheckHasCoinUrl, request)
	if err != nil {
		return VideoHasCoinsResponse{}, err
	}
	response := VideoHasCoinsResponse{}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return VideoHasCoinsResponse{}, err
	}
	return response, nil
}

// CheckVideoHasCoinsByAid checks if the logged-in user has sent coins to this video specified by the Aid
func (c *Client) CheckVideoHasCoinsByAid(aid int) (VideoHasCoinsResponse, error) {
	return c.checkVideoHasCoins(VideoRequestAid{Aid: aid})
}

// CheckVideoHasCoinsByBvid checks if the logged-in user has sent coins to this video specified by the Bvid
func (c *Client) CheckVideoHasCoinsByBvid(bvid string) (VideoHasCoinsResponse, error) {
	return c.checkVideoHasCoins(VideoRequestBvid{Bvid: bvid})
}

//
// Change (add/remove) video to/from the favorite list (收藏)
//

// VideoChangeFavRequest is the request payload of changing the favorite status of the video specified by Aid
type VideoChangeFavRequest struct {
	Aid         int    `url:"rid"`
	Type        int    `url:"type"`
	AddMediaIds []int  `url:"add_media_ids"`
	DelMediaIds []int  `url:"del_media_ids"`
	Csrf        string `url:"csrf"`
	Jsonp       string `url:"jsonp"`
}

// VideoChangeFavResponseCode is the status code of the response for changing the favorite status of the video
type VideoChangeFavResponseCode int

const (
	VideoChangeFavSuccess        VideoChangeFavResponseCode = 0
	VideoChangeFavNotLoggedIn    VideoChangeFavResponseCode = -101
	VideoChangeFavWrongCsrf      VideoChangeFavResponseCode = -111
	VideoChangeFavBadRequest     VideoChangeFavResponseCode = -400
	VideoChangeFavNoPermission   VideoChangeFavResponseCode = -403
	VideoChangeFavVideoNotExists VideoChangeFavResponseCode = 10003
	VideoChangeFavAlreadyAdded   VideoChangeFavResponseCode = 11201
	VideoChangeFavReachLimit     VideoChangeFavResponseCode = 11203
	VideoChangeFavIncorrectArgs  VideoChangeFavResponseCode = 72010017
)

// VideoChangeFavResponse is the response for changing the favorite status of the video
type VideoChangeFavResponse struct {
	Code    VideoChangeFavResponseCode `json:"code"`
	Message string                     `json:"message"`
	Data    struct {
		FavByNonFollower bool `json:"prompt"`
	}
}

// changeVideoFav changes the favorite status of the video
func (c *Client) changeVideoFav(request interface{}) (VideoChangeFavResponse, error) {
	responseBody, _, err := HttpPostWithParamsReferer(c.client, c.config.Endpoints.VideoChangeFavUrl, request, "https://www.bilibili.com/")
	if err != nil {
		return VideoChangeFavResponse{}, err
	}
	response := VideoChangeFavResponse{}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return VideoChangeFavResponse{}, err
	}
	return response, nil
}

// ChangeVideoFavByAid changes the favorite status of the video specified by Aid
func (c *Client) ChangeVideoFavByAid(aid int, addMedias []int, delMedias []int) (VideoChangeFavResponse, error) {
	csrf, err := c.getCookieValueByName("bili_jct")
	if err != nil {
		return VideoChangeFavResponse{}, err
	}
	return c.changeVideoFav(VideoChangeFavRequest{
		Aid:         aid,
		Type:        2,
		AddMediaIds: addMedias,
		DelMediaIds: delMedias,
		Csrf:        csrf,
		Jsonp:       "jsonp",
	})
}

//
// Check if the video has been favored or not
//

// VideoFavoredResponse is the response for checking if a video has been favored by the logged-in user
type VideoFavoredResponse struct {
	Code    VideoChangeFavResponseCode `json:"code"`
	Message string                     `json:"message"`
	Ttl     int                        `json:"ttl"`
	Data    struct {
		Favoured bool `json:"favoured"`
	} `json:"data"`
}

// checkVideoFavored checks if a video has been favored by the logged-in user
func (c *Client) checkVideoFavored(request interface{}) (VideoFavoredResponse, error) {
	responseBody, err := HttpGetWithParams(c.client, c.config.Endpoints.VideoCheckHasCoinUrl, request)
	if err != nil {
		return VideoFavoredResponse{}, err
	}
	response := VideoFavoredResponse{}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return VideoFavoredResponse{}, err
	}
	return response, nil
}

// CheckVideoFavoredByAid checks if a video has been favored by the logged-in user specified by Aid
func (c *Client) CheckVideoFavoredByAid(aid int) (VideoFavoredResponse, error) {
	return c.checkVideoFavored(VideoRequestAid{Aid: aid})
}

// CheckVideoFavoredByBvid checks if a video has been favored by the logged-in user specified by Bvid
func (c *Client) CheckVideoFavoredByBvid(bvid string) (VideoFavoredResponse, error) {
	return c.checkVideoFavored(VideoRequestBvid{Bvid: bvid})
}
