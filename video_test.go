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
