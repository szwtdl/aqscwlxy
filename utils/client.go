package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type HttpClient struct {
	client  *http.Client
	domain  string
	cookies map[string]string
}

func NewClient() *HttpClient {
	return &HttpClient{
		client:  &http.Client{},
		domain:  "https://gd.aqscwlxy.com/gd_api",
		cookies: make(map[string]string),
	}
}

func (h *HttpClient) DoPost(postUrl string, headers map[string]string, postData map[string]string) ([]byte, error) {
	var data []byte
	var err error
	contentType, exists := headers["Content-Type"]
	if !exists {
		contentType = "application/json"
	}
	switch contentType {
	case "application/json":
		data, err = json.Marshal(postData)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal JSON: %w", err)
		}
	case "application/x-www-form-urlencoded":
		postDataValues := make(url.Values)
		for k, v := range postData {
			postDataValues.Set(k, ToString(v))
		}
		data = []byte(postDataValues.Encode())
	default:
		return nil, fmt.Errorf("unsupported Content-Type: %s", contentType)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", h.domain, postUrl), bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return h.doRequest(req)
}

func (h *HttpClient) DoGet(url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", h.domain, url), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return h.doRequest(req)
}

func (h *HttpClient) doRequest(req *http.Request) ([]byte, error) {
	res, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	return body, nil
}
