package http

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	ContentType string
}

var client *Client

type IHttpClient interface {
	Load()
	Get(url string) *http.Response
	Post(url string) *http.Response
}

func (h *Client) Load() {
	httpClient := &Client{
		ContentType: "application/json",
	}

	client = httpClient
}

func (h *Client) Get(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	return resp
}

func (h *Client) Post(url string, body any) (*http.Response, error) {
	jsonBody, _ := json.Marshal(body)
	bodyBuffer := bytes.NewBuffer(jsonBody)

	resp, err := http.Post(url, client.ContentType, bodyBuffer)
	if err != nil {
		panic(err)
	}
	return resp, err
}
