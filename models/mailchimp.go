package models

import (
	"encoding/json"
	"errors"
	"github.com/byrnedo/tictail_to_mailchimp_forwarder/config"
	"io/ioutil"
	"net/http"
)

type Mailchimp struct {
	url    string
	apiKey string
}

type List struct {
	Id   string
	Name string
}

func NewMailchimp(c *config.Config) *Mailchimp {
	return &Mailchimp{
		c.Mailchimp.Url,
		c.Mailchimp.AccessToken,
	}
}

func (m *Mailchimp) GetLists() (lists *[]List, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", m.url, nil)
	if err != nil {
		return lists, errors.New("Failed to prepare GET request: " + err.Error())
	}

	req.Header.Add("Authorization", "apiKey "+m.apiKey)

	response, err := client.Do(req)
	if err != nil {
		return lists, errors.New("Failed to do GET request: " + err.Error())
	}

	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return lists, errors.New("Failed to read request body: " + err.Error())
	}
	*lists = make([]List, 0, 0)

	if err = json.Unmarshal(contents, lists); err != nil {
		return lists, errors.New("Error unmarshalling data: " + err.Error())
	}
	return lists, nil
}
