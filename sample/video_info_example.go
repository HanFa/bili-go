package main

import (
	"bili"
	"log"
)

func main() {
	client, err := bili.New("./config.json")
	if err != nil {
		log.Fatalf("cannot new a bili client\n")
	}
	response, err := client.GetVideoInfoByAid(85440373)
	if err != nil {
		log.Fatalf("cannot get video info: %s\n", err.Error())
	}
	log.Printf("response: %v\n", response)
}
