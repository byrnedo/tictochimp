package models

import (
	"encoding/json"
	"errors"
	"github.com/byrnedo/tictochimp/utils"
)

type Mailchimp struct {
	url    string
	apiKey string
	client *utils.RestClient
}

const API_VERSION = "3.0"

type List struct {
	Id   string
	Name string
}

type Subscriber struct {
	Email     string
	FirstName string
	LastName  string
}

func NewMailchimp(url string, apiKey string) *Mailchimp {

	client := utils.NewRestClient()

	client.Headers = map[string]string{
		"Authorization": "apiKey " + apiKey,
	}
	return &Mailchimp{
		url + "/" + API_VERSION,
		apiKey,
		client,
	}
}

func (m *Mailchimp) GetLists() ([]List, error) {

	responseData := struct {
		Lists []List
	}{}

	err := m.client.Get(m.url + "/lists")
	if err != nil {
		return nil, errors.New("Failed to do GET request: " + err.Error())
	}

	status := m.client.ResponseStatus()
	if status != "200" && status != "201" {
		return nil, errors.New("Got non-success status code: " + status)
	}

	err = json.Unmarshal(m.client.ResponseBody(), &responseData)

	return responseData.Lists, nil
}

type mergeFields struct {
	FirstName string `json:"FNAME"`
	LastName  string `json:"LNAME"`
}
type memberRequest struct {
	EmailType    string      `json:"email_type"`
	EmailAddress string      `json:"email_address"`
	StatusIfNew  string      `json:"status_if_new"`
	Status       string      `json:"status"`
	MergeFields  mergeFields `json:"merge_fields"`
	Language     string      `json:"language"`
	Vip          bool        `json:"vip"`
}

func newMemeberRequest(subscriber Subscriber) *memberRequest {
	return &memberRequest{
		EmailType:    "html",
		EmailAddress: subscriber.Email,
		StatusIfNew:  "subscribed",
		Status:       "subscribed",
		MergeFields:  mergeFields{subscriber.FirstName, subscriber.LastName},
		Language:     "sv",
		Vip:          false,
	}
}

func (m *Mailchimp) AddSubscriber(sub Subscriber, listId string) error {
	err := m.client.Post(m.url+"/lists/"+listId+"/members", newMemeberRequest(sub))
	if err == nil {
		status := m.client.ResponseStatus()
		if status != "200" && status != "201" {
			return errors.New("Got non-success status code: " + status)
		}
	}

	return err
}
