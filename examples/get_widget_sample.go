package main

import (
	"github.com/hanfa/bili-go"
	"log"
)

func main() {
	_, err := bili.New()
	if err != nil {
		log.Fatalf("Cannot new a bili-client: " + err.Error())
	}
}
