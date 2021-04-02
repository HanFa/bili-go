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
	err = client.DownloadByAid(bili.DownloadOptionAid{
		Aid: 671597785,
		DownloadOptionCommon: bili.DownloadOptionCommon{
			Page:       0,
			Resolution: bili.Stream1080P,
			Mode:       bili.StreamFlv,
			Allow4K:    true,
			OutPath:    "/tmp/out.flv",
		},
	}, true)

	if err != nil {
		log.Fatalln(err.Error())
	}
}
