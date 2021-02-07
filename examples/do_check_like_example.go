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
	resp, err := client.Login("13262272815", "han5fang1")
	if err != nil {
		log.Fatalln("Error when doing login with password " + err.Error())
	}
	if resp.Code != bili.LoginSuccess {
		log.Fatalln("Cannot login: " + resp.Message)
	}
	_, err = client.LikeVideoByBvid("BV1bU4y1x7A1", bili.VideoLike) // like
	if err != nil {
		log.Fatalln("Error when liking: " + err.Error())
	}
	checkLikeResponse, err := client.CheckVideoLikeByBvid("BV1bU4y1x7A1") // check like
	if err != nil {
		log.Fatalln("Error when check like status: " + err.Error())
	}
	if checkLikeResponse.Data == bili.Liked {
		log.Println("As expected, you have liked this video")
	} else {
		log.Fatalln("What??")
	}
}
