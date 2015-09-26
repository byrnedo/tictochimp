package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type RestClient struct {
	client           *http.Client
	Headers          map[string]string
	lastResponse     http.Response
	lastResponseBody []byte
}

func NewRestClient() *RestClient {
	return &RestClient{&http.Client{}, nil, http.Response{}, nil}
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

	os.Stdout.WriteString("Making request to " + url + "\n")

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

	os.Stdout.WriteString("Making request to " + url + "\n")

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
