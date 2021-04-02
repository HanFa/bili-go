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
	err = client.DoCaptcha()
	if err != nil {
		log.Fatalln("Error when doing captcha " + err.Error())
	}
}
