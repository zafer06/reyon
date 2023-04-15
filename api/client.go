package main

import (
	"encoding/json"
    "io"
    "log"
    "net/http"
)

type TopicResponse struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Error string `json:"error"`
	Data []Topic `json:"data"`
}

type Topic struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	User        string `json:"user"`
	Date        string `json:"date"`
	Reply       string `json:"reply"`
	Description string `json:"description"`
}

func GetTopicList() (TopicResponse, error) {
	data, err := sendRequest("http://bot:4003/list")
	if err != nil {
		return TopicResponse{}, err
	}

	var topicResponse TopicResponse
	err = json.Unmarshal([]byte(data), &topicResponse)
	if err != nil {
		return TopicResponse{}, err
	}

	return topicResponse, nil
}

func sendRequest(url string) (string, error) {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer ffc63b1afa1de95856e5117f829d9b3d612551ac")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if (err != nil) {
		log.Println(err)
	}

	return string(body), err
}