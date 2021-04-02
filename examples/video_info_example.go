package main

import (
	"github.com/hanfa/bili-go"
	"log"
)

func main() {
	client, err := bili.New()
	if err != nil {
		log.Fatalf("cannot new a bili client\n")
	}
	response, err := client.GetVideoInfoByAid(85440373)
	if err != nil {
		log.Fatalf("cannot get video info: %s\n", err.Error())
	}
	log.Printf("response: %v\n", response)
}
