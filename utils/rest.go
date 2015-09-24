package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type RestClient struct {
	client  *http.Client
	Headers map[string]string
}

func NewRestClient() *RestClient {
	return &RestClient{&http.Client{}, nil}
}

func (c *RestClient) Get(url string, responseStruct interface{}) error {
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

	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(contents, responseStruct); err != nil {
		return err
	}
	return nil
}

func (c *RestClient) Post(url string, data interface{}, responseStruct interface{}) error {

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

	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(contents, responseStruct); err != nil {
		return err
	}
	return nil
}
