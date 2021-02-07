package main

import (
	bili "bili"
	"log"
)

func main() {
	_, err := bili.New("config.json")
	if err != nil {
		log.Fatalf("Cannot new a bili-client: " + err.Error())
	}
}
