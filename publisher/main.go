package main

import (
	"log"
	"os"

	"github.com/ququzone/vote-sbt/publisher/fetcher"
	"github.com/ququzone/vote-sbt/publisher/pusher"
)

func main() {
	token := os.Getenv("TOKEN")
	var skip uint64 = 0
	for {
		voters, err := fetcher.Fetch("0x8e171d0b0d2d74e82dc37359411d9eb286482a6c08ad52563ca265d84de738f9", skip)
		if err != nil {
			log.Fatalf("fetch voter error: %v\n", err)
		}
		for _, v := range voters {
			err := pusher.Push(token, v[2:], "4")
			if err != nil {
				log.Printf("post mint data to node error: %v\n", err)
			}
		}
		if len(voters) != 10 {
			break
		}
		skip += 10
	}
}
