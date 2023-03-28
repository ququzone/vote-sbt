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
		voters, err := fetcher.Fetch("0x2f97ac5cb8467fe16d4dd0b52d95aeba67e1196df2858baa710b238768a70b78", skip)
		if err != nil {
			log.Fatalf("fetch voter error: %v\n", err)
		}
		for _, v := range voters {
			err := pusher.Push(token, v[2:], "1")
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
