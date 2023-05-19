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
		voters, err := fetcher.Fetch("0x3243e5577acb0b8b2f4d9728071d713b02a26b1c247dbbd727e0b3a1b9343659", skip)
		if err != nil {
			log.Fatalf("fetch voter error: %v\n", err)
		}
		for _, v := range voters {
			err := pusher.Push(token, v[2:], "3")
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
