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

type VideoState int
type VideoDimension struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Rotate int `json:"rotate"`
}

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

//
//Check a video has received coins from you or not
//

type VideoHasCoinsResponse struct {
	Code    VideoAddCoinResponseCode `json:"code"`
	Message string                   `json:"message"`
	Ttl     int                      `json:"ttl"`
	Data    struct {
		Multiply int `json:"multiply"`
	} `json:"data"`
}

func (c *Client) checkVideoHasCoins(request interface{}) (VideoHasCoinsResponse, error) {
	responseBody, err := HttpGetWithParams(c.client, c.Config.Endpoints.VideoCheckHasCoinUrl, request)
	if err != nil {
		return VideoHasCoinsResponse{}, err
	}
	response := VideoHasCoinsResponse{}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return VideoHasCoinsResponse{}, err
	}
	return response, nil
}

func (c *Client) CheckVideoHasCoinsByAid(aid int) (VideoHasCoinsResponse, error) {
	return c.checkVideoHasCoins(VideoRequestAid{Aid: aid})
}

func (c *Client) CheckVideoHasCoinsByBvid(bvid string) (VideoHasCoinsResponse, error) {
	return c.checkVideoHasCoins(VideoRequestBvid{Bvid: bvid})
}

//
//Change (add/remove) video to/from the favorite list
//

type VideoChangeFavRequest struct {
	Aid         int    `url:"rid"`
	Type        int    `url:"type"`
	AddMediaIds []int  `url:"add_media_ids"`
	DelMediaIds []int  `url:"del_media_ids"`
	Csrf        string `url:"csrf"`
	Jsonp       string `url:"jsonp"`
}

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

type VideoChangeFavResponse struct {
	Code    VideoChangeFavResponseCode `json:"code"`
	Message string                     `json:"message"`
	Data    struct {
		FavByNonFollower bool `json:"prompt"`
	}
}

func (c *Client) changeVideoFav(request interface{}) (VideoChangeFavResponse, error) {
	responseBody, _, err := HttpPostWithParamsReferer(c.client, c.Config.Endpoints.VideoChangeFavUrl, request, "https://www.bilibili.com/")
	if err != nil {
		return VideoChangeFavResponse{}, err
	}
	response := VideoChangeFavResponse{}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return VideoChangeFavResponse{}, err
	}
	return response, nil
}

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

type VideoFavoredResponse struct {
	Code    VideoChangeFavResponseCode `json:"code"`
	Message string                     `json:"message"`
	Ttl     int                        `json:"ttl"`
	Data    struct {
		Favoured bool `json:"favoured"`
	} `json:"data"`
}

func (c *Client) checkVideoFavored(request interface{}) (VideoFavoredResponse, error) {
	responseBody, err := HttpGetWithParams(c.client, c.Config.Endpoints.VideoCheckHasCoinUrl, request)
	if err != nil {
		return VideoFavoredResponse{}, err
	}
	response := VideoFavoredResponse{}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return VideoFavoredResponse{}, err
	}
	return response, nil
}

func (c *Client) CheckVideoFavoredByAid(aid int) (VideoFavoredResponse, error) {
	return c.checkVideoFavored(VideoRequestAid{Aid: aid})
}

func (c *Client) CheckVideoFavoredByBvid(bvid string) (VideoFavoredResponse, error) {
	return c.checkVideoFavored(VideoRequestBvid{Bvid: bvid})
}
