package mailchimp

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/byrnedo/tictochimp/models/mailchimp/spec"
	"github.com/byrnedo/tictochimp/utils"
	"strings"
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

func NewMailchimp(apiKey string) (*Mailchimp, error) {

	client := utils.NewRestClient()

	client.Headers = map[string]string{
		"Authorization": "apiKey " + apiKey,
	}

	keyParts := strings.Split(apiKey, "-")
	if len(keyParts) != 2 {
		return nil, errors.New("Access Token must end with '-<datacenter code>', eg. '-us10'")
	}
	dataCenter := keyParts[1]
	return &Mailchimp{
		"https://" + dataCenter + ".api.mailchimp.com/" + API_VERSION,
		apiKey,
		client,
	}, nil
}

func (m *Mailchimp) GetLists() ([]spec.List, error) {

	responseData := spec.ListsResponse{}
	err := m.client.Get(m.url + "/lists")
	if err != nil {
		return nil, errors.New("Failed to do GET request: " + err.Error())
	}

	status := m.client.ResponseStatus()
	if status != 200 && status != 201 {
		return nil, errors.New("Got non-success status code: " + string(status))
	}

	err = json.Unmarshal(m.client.ResponseBody(), &responseData)

	return responseData.Lists, err
}

func (m *Mailchimp) Get10ListMembers(offset int, listID string) (data []spec.Member, total int, err error) {

	responseData := spec.MembersResponse{}
	err = m.client.Get(fmt.Sprintf("%s/lists/%s/members?offset=%d", m.url, listID, offset))
	if err != nil {
		return nil, 0, errors.New("Failed to do GET request: " + err.Error())
	}

	status := m.client.ResponseStatus()
	if status != 200 && status != 201 {
		return nil, 0, errors.New("Got non-success status code: " + string(status))
	}

	err = json.Unmarshal(m.client.ResponseBody(), &responseData)
	return responseData.Members, responseData.TotalItems, err
}

func (m *Mailchimp) GetAllListMembers(listID string) (totalList []spec.Member, err error) {
	totalList = make([]spec.Member, 0, 0)

	offset := 0
	keepGoing := true

	for keepGoing {
		batchResult, total, err := m.Get10ListMembers(offset, listID)
		if err != nil {
			return nil, errors.New("Failed to get batch of 10 subscribers from list: " + err.Error())
		}
		currentBatchSize := len(batchResult)
		offset = offset + currentBatchSize
		if offset >= total {
			keepGoing = false
		}
		totalList = append(totalList, batchResult...)
	}
	return totalList, nil
}

func newMemeberRequest(subscriber Subscriber) *spec.MemberRequest {
	return &spec.MemberRequest{
		EmailType:    "html",
		EmailAddress: subscriber.Email,
		StatusIfNew:  "subscribed",
		Status:       "subscribed",
		MergeFields:  &spec.MergeFields{subscriber.FirstName, subscriber.LastName},
		Language:     "sv",
		Vip:          false,
	}
}

func (m *Mailchimp) AddSubscriber(sub Subscriber, listId string) error {
	err := m.client.Post(m.url+"/lists/"+listId+"/members", newMemeberRequest(sub))
	if err == nil {
		status := m.client.ResponseStatus()
		if status != 200 && status != 201 {
			return errors.New("Got non-success status code: " + string(status))
		}
	}

	return err
}
