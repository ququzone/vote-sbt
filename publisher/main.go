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
		voters, err := fetcher.Fetch("0x38e56cff9f04b88c29447a6b71734274f65931cc8af049f82a268abb7474af0c", skip)
		if err != nil {
			log.Fatalf("fetch voter error: %v\n", err)
		}
		for _, v := range voters {
			err := pusher.Push(token, v[2:], "2")
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
