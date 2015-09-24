package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type RestClient struct {
	client       *http.Client
	Headers      map[string]string
	lastResponse *http.Response
}

func NewRestClient() *RestClient {
	return &RestClient{&http.Client{}, nil, nil}
}

func (c *RestClient) ResponseBody() []byte {
	data, _ := ioutil.ReadAll(c.lastResponse.Body)
	return data
}

func (c *RestClient) ResponseStatus() string {
	return c.lastResponse.Status
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
	c.lastResponse = response
	if err != nil {
		return err
	}

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
	c.lastResponse = response
	if err != nil {
		return err
	}

	defer response.Body.Close()

	return nil
}
