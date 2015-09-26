package models

import (
	"encoding/json"
	"errors"
	"github.com/byrnedo/tictochimp/models/mailchimpSpec"
	"github.com/byrnedo/tictochimp/utils"
)

type Mailchimp struct {
	url    string
	apiKey string
	client *utils.RestClient
}

const API_VERSION = "3.0"

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

func (m *Mailchimp) GetLists() ([]mailchimpSpec.List, error) {

	responseData := struct {
		Lists []mailchimpSpec.List
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

	return responseData.Lists, err
}

func newMemeberRequest(subscriber Subscriber) *mailchimpSpec.MemberRequest {
	return &mailchimpSpec.MemberRequest{
		EmailType:    "html",
		EmailAddress: subscriber.Email,
		StatusIfNew:  "subscribed",
		Status:       "subscribed",
		MergeFields:  &mailchimpSpec.MergeFields{subscriber.FirstName, subscriber.LastName},
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
