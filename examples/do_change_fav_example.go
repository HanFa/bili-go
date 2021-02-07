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
	changeFavResponse, err := client.ChangeVideoFavByAid(671597785, []int{180320832}, []int{})
	if err != nil {
		log.Fatalln("Error when changing fav: " + err.Error())
	}
	log.Println(changeFavResponse)
}
