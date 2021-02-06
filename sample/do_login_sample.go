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
	err = client.Login("13262272815", "han5fang1")
	if err != nil {
		log.Fatalln("Error when doing login with password " + err.Error())
	}
}
