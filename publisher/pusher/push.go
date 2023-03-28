package pusher

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Body struct {
	Events []Event `json:"events"`
}

type Header struct {
	PubId     string `json:"pub_id"`
	PubTime   int64  `json:"pub_time"`
	EventType string `json:"event_type"`
	Token     string `json:"token"`
}

type Event struct {
	Header  Header `json:"header"`
	Payload string `json:"payload"`
}

func Push(token string, account string, tokenId string) error {
	url := "http://34.146.117.200:3001/api/w3bapp/event/vote_sbt_testnet"
	method := "POST"

	payload := fmt.Sprintf(`{"payload": {"account": "%s","tokenId": "%s"}}`, account, tokenId)
	payload = base64.StdEncoding.EncodeToString([]byte(payload))
	body := &Body{
		Events: []Event{{
			Header: Header{
				PubId:     "vote_sbt_publisher",
				PubTime:   time.Now().Unix(),
				EventType: "ANY",
				Token:     token,
			},
			Payload: payload,
		}},
	}
	bodyData, _ := json.Marshal(body)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyData))

	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
