package bili

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_GetVideoInfoByAid(t *testing.T) {
	response, err := client.GetVideoInfoByAid(85440373)
	assert.Nil(t, err, "cannot get video by aid")
	assert.Equal(t, response.Code, VideoSuccess, "response doesn't have success code")
	assert.Equal(t, response.Data.Title, "当我给拜年祭的快板加了电音配乐…")
}

func TestClient_GetVideoInfoByBvid(t *testing.T) {
	response, err := client.GetVideoInfoByBvid("BV117411r7R1")
	assert.Nil(t, err, "cannot get video by bvid")
	assert.Equal(t, response.Code, VideoSuccess, "response doesn't have success code")
	assert.Equal(t, response.Data.Title, "当我给拜年祭的快板加了电音配乐…")
}

func TestClient_GetVideoDescriptionByAid(t *testing.T) {
	response, err := client.GetVideoDescriptionByAid(85440373)
	assert.Nil(t, err, "cannot get video desc by aid")
	assert.Equal(t, response.Data, "【CB想说的】看完拜年祭之后最爱的一个节目！给有快板的部分简单加了一些不同风格的配乐hhh，感谢沃玛画的我！太可爱了哈哈哈哈哈哈哈！！！\n【Warma想说的】我画了打碟的CB，画风为了还原原版视频所以参考了四迹老师的画风，四迹老师的画真的太可爱啦！不过其实在画的过程中我遇到了一个问题，CB的耳机……到底是戴在哪个耳朵上呢？\n\n原版：av78977080\n编曲（配乐）：Crazy Bucket\n人声（配音）：Warma/谢拉\n曲绘：四迹/Warma\n动画：四迹/Crazy Bucket\n剧本：Mokurei-木灵君\n音频后期：DMYoung/纳兰寻风/Crazy Bucket\n包装：破晓天")
}

func TestClient_GetVideoDescriptionByBvid(t *testing.T) {
	response, err := client.GetVideoDescriptionByBvid("BV117411r7R1")
	assert.Nil(t, err, "cannot get video desc by bvid")
	assert.Equal(t, response.Data, "【CB想说的】看完拜年祭之后最爱的一个节目！给有快板的部分简单加了一些不同风格的配乐hhh，感谢沃玛画的我！太可爱了哈哈哈哈哈哈哈！！！\n【Warma想说的】我画了打碟的CB，画风为了还原原版视频所以参考了四迹老师的画风，四迹老师的画真的太可爱啦！不过其实在画的过程中我遇到了一个问题，CB的耳机……到底是戴在哪个耳朵上呢？\n\n原版：av78977080\n编曲（配乐）：Crazy Bucket\n人声（配音）：Warma/谢拉\n曲绘：四迹/Warma\n动画：四迹/Crazy Bucket\n剧本：Mokurei-木灵君\n音频后期：DMYoung/纳兰寻风/Crazy Bucket\n包装：破晓天")
}

func TestClient_LikeVideoByAid_NotLoggedIn(t *testing.T) {
	response, err := client.LikeVideoByAid(82887239, VideoLike)
	assert.Nil(t, err, "cannot like video")
	assert.Equal(t, response.Code, VideoLikeNotLoggedIn)
}

func TestClient_LikeVideoByBvid_NotLoggedIn(t *testing.T) {
	response, err := client.LikeVideoByBvid("BV1bU4y1x7A1", VideoLike)
	assert.Nil(t, err, "cannot like video")
	assert.Equal(t, response.Code, VideoLikeNotLoggedIn)
}

func TestClient_CheckVideoLikeByAid_NotLoggedIn(t *testing.T) {
	response, err := client.CheckVideoLikeByAid(82887239)
	assert.Nil(t, err, "cannot check video like")
	assert.Equal(t, response.Code, VideoLikeNotLoggedIn)
}

func TestClient_CheckVideoLikeByBvid_NotLoggedIn(t *testing.T) {
	response, err := client.CheckVideoLikeByBvid("BV1bU4y1x7A1")
	assert.Nil(t, err, "cannot check video like")
	assert.Equal(t, response.Code, VideoLikeNotLoggedIn)
}

func TestClient_AddCoinToVideoByAid_NotLoggedIn(t *testing.T) {
	response, err := client.AddCoinToVideoByAid(82887239, 1, true)
	assert.Nil(t, err, "cannot add coins to video")
	assert.Equal(t, response.Code, VideoAddCoinNotLoggedIn)
}

func TestClient_AddCoinToVideoByBvid_NotLoggedIn(t *testing.T) {
	response, err := client.AddCoinToVideoByBvid("BV1bU4y1x7A1", 1, true)
	assert.Nil(t, err, "cannot add coins to video")
	assert.Equal(t, response.Code, VideoAddCoinNotLoggedIn)
}

func TestClient_CheckVideoHasCoinsByAid_NotLoggedIn(t *testing.T) {
	response, err := client.CheckVideoHasCoinsByAid(82887239)
	assert.Nil(t, err, "cannot check video like")
	assert.Equal(t, response.Code, VideoAddCoinNotLoggedIn)
}

func TestClient_CheckVideoHasCoinsByBvid_NotLoggedIn(t *testing.T) {
	response, err := client.CheckVideoHasCoinsByBvid("BV1bU4y1x7A1")
	assert.Nil(t, err, "cannot check video like")
	assert.Equal(t, response.Code, VideoAddCoinNotLoggedIn)
}

func TestClient_ChangeVideoFavByAid_NotLoggedIn(t *testing.T) {
	response, err := client.ChangeVideoFavByAid(671597785, []int{180320832}, []int{})
	assert.Nil(t, err, "cannot change fav for video")
	assert.Equal(t, response.Code, VideoChangeFavNotLoggedIn)
}

func TestClient_CheckVideoFavoredByAid_NotLoggedIn(t *testing.T) {
	response, err := client.CheckVideoFavoredByAid(671597785)
	assert.Nil(t, err, "cannot check video fav")
	assert.Equal(t, response.Code, VideoChangeFavNotLoggedIn)
}

func TestClient_CheckVideoFavoredByBvid_NotLoggedIn(t *testing.T) {
	response, err := client.CheckVideoFavoredByBvid("BV1bU4y1x7A1")
	assert.Nil(t, err, "cannot check video fav")
	assert.Equal(t, response.Code, VideoChangeFavNotLoggedIn)
}
