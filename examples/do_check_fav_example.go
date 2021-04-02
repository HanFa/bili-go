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
	checkFavResp, err := client.CheckVideoFavoredByBvid("BV1bU4y1x7A1") // check fav
	if err != nil {
		log.Fatalln("Error when checking favoured or not")
	}
	log.Printf("Favored: %d\n", checkFavResp.Data.Favoured)
}
