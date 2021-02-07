package main

import (
	"bili"
	"log"
)

func main() {
	client, err := bili.New("config.json")
	if err != nil {
		log.Fatalln("Cannot new a bili-client: " + err.Error())
	}
	resp, err := client.Login("username", "password") // @TODO: Change me
	if err != nil {
		log.Fatalln("Error when doing login with password " + err.Error())
	}
	if resp.Code != bili.LoginSuccess {
		log.Fatalln("Cannot login: " + resp.Message)
	}
	likeResp, err := client.LikeVideoByBvid("BV1bU4y1x7A1", bili.VideoLike) // like
	if err != nil {
		log.Fatalln("Error when liking: " + err.Error())
	}
	log.Println(likeResp)
}
