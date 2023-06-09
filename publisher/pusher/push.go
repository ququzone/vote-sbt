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
	url := "http://35.221.109.71:8888/srv-applet-mgr/v0/event/voted_sbt"
	method := "POST"

	payload := fmt.Sprintf(`{"account": "%s","tokenId": "%s"}`, account, tokenId)
	payload = base64.StdEncoding.EncodeToString([]byte(payload))
	body := &Body{
		Events: []Event{{
			Header: Header{
				PubId:     "vote_sbt_publisher",
				PubTime:   time.Now().Unix() * 1000,
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

	// resBody, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(string(resBody))
	return nil
}
