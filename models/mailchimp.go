package models

import (
	"errors"
	"github.com/byrnedo/tictochimp/utils"
)

type Mailchimp struct {
	url    string
	apiKey string
}

const API_VERSION = "3.0"

type List struct {
	Id   string
	Name string
}

func NewMailchimp(url string, apiKey string) *Mailchimp {
	return &Mailchimp{
		url + "/" + API_VERSION,
		apiKey,
	}
}

func (m *Mailchimp) GetLists() ([]List, error) {
	client := utils.NewRestClient()

	client.Headers = map[string]string{
		"Authorization": "apiKey " + m.apiKey,
	}

	responseData := struct {
		Lists []List
	}{}

	err := client.Get(m.url+"/lists", &responseData)
	if err != nil {
		return nil, errors.New("Failed to do GET request: " + err.Error())
	}

	return responseData.Lists, nil
}
