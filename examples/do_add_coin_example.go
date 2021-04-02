package main

import (
	"github.com/hanfa/bili-go"
	"log"
)

func main() {
	client, err := bili.New()
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
	addCoinResp, err := client.AddCoinToVideoByBvid("BV1bU4y1x7A1", 1, true)
	if err != nil {
		log.Fatalln("Error when adding coin: " + err.Error())
	}
	log.Println(addCoinResp)
}
