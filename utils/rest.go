package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type RestClient struct {
	client           *http.Client
	Headers          map[string]string
	lastResponse     http.Response
	lastResponseBody []byte
}

func NewRestClient() *RestClient {
	timeout := time.Duration(5 * time.Second)
	return &RestClient{&http.Client{Timeout: timeout}, nil, http.Response{}, nil}
}

func (c *RestClient) ResponseBody() []byte {
	return c.lastResponseBody
}

func (c *RestClient) ResponseStatus() int {
	return c.lastResponse.StatusCode
}

func (c *RestClient) Get(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	for name, value := range c.Headers {
		req.Header.Add(name, value)
	}

	response, err := c.client.Do(req)
	if err != nil {
		return err
	}
	c.lastResponse = *response
	c.lastResponseBody, _ = ioutil.ReadAll(response.Body)

	defer response.Body.Close()

	return nil
}

func (c *RestClient) Post(url string, data interface{}) error {

	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(byteData))
	if err != nil {
		return err
	}

	for name, value := range c.Headers {
		req.Header.Add(name, value)
	}

	response, err := c.client.Do(req)
	if err != nil {
		return err
	}
	c.lastResponse = *response
	c.lastResponseBody, _ = ioutil.ReadAll(response.Body)

	defer response.Body.Close()

	return nil
}
