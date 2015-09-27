package mailchimp

import (
	"encoding/json"
	"github.com/byrnedo/tictochimp/models/mailchimp/spec"
	"github.com/jarcoal/httpmock"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

const (
	MAILCHIMP_TEST_URL = "https://us10.api.mailchimp.com"
	MAILCHIMP_TEST_KEY = "testToken1234-us10"
)

func TestGetLists(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", MAILCHIMP_TEST_URL+"/3.0/lists",
		func(req *http.Request) (*http.Response, error) {
			if req.Header.Get("Authorization") == "apiKey "+MAILCHIMP_TEST_KEY {
				return httpmock.NewStringResponse(200, MAILCHIMP_MOCK_GET_LISTS_200_RESPONSE), nil
			}
			return httpmock.NewStringResponse(401, "{}"), nil
		},
	)

	mc := NewMailchimp(MAILCHIMP_TEST_KEY)

	data, err := mc.GetLists()
	if err != nil {
		t.Error(err.Error())
	}

	expectedData := spec.ListsResponse{}
	err = json.Unmarshal([]byte(MAILCHIMP_MOCK_GET_LISTS_200_RESPONSE), &expectedData)
	if err != nil {
		t.Error(err.Error())
	}

	if !reflect.DeepEqual(expectedData.Lists, data) {
		t.Error("Response and mock data didn't match")
	}

}

func TestGet10ListMembers(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", MAILCHIMP_TEST_URL+"/3.0/lists/x1234/members?offset=0",
		func(req *http.Request) (*http.Response, error) {
			if req.Header.Get("Authorization") == "apiKey "+MAILCHIMP_TEST_KEY {
				return httpmock.NewStringResponse(200, MAILCHIMP_MOCK_GET_LIST_MEMBERS_200_RESPONSE), nil
			}
			return httpmock.NewStringResponse(401, "{}"), nil
		},
	)

	mc := NewMailchimp(MAILCHIMP_TEST_KEY)

	data, total, err := mc.Get10ListMembers(0, "x1234")
	if err != nil {
		t.Error(err.Error())
	}

	if total != 1 {
		t.Error("Incorrect total found")
	}

	expectedData := spec.MembersResponse{}
	err = json.Unmarshal([]byte(MAILCHIMP_MOCK_GET_LIST_MEMBERS_200_RESPONSE), &expectedData)
	if err != nil {
		t.Error(err.Error())
	}

	if !reflect.DeepEqual(expectedData.Members, data) {
		t.Error("Response and mock data didn't match")
	}

}

func TestAddSubscriberToList(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", MAILCHIMP_TEST_URL+"/3.0/lists/x1234/members",
		func(req *http.Request) (*http.Response, error) {
			if req.Header.Get("Authorization") == "apiKey "+MAILCHIMP_TEST_KEY {
				bodyBytes, _ := ioutil.ReadAll(req.Body)
				expectedData := map[string]interface{}{
					"email_type":    "html",
					"email_address": "donal@test.com",
					"status":        "subscribed",
					"status_if_new": "subscribed",
					"merge_fields": map[string]interface{}{
						"FNAME": "Donal",
						"LNAME": "Test",
					},
					"language": "sv",
					"vip":      false,
				}

				var receivedMap map[string]interface{}
				if err := json.Unmarshal(bodyBytes, &receivedMap); err != nil {
					return httpmock.NewStringResponse(500, `{"error","Failed to unmarshal body to bytes"}`), nil
				}

				if !reflect.DeepEqual(expectedData, receivedMap) {
					t.Logf("Expected: %+v", expectedData)
					t.Logf("Received: %+v", receivedMap)
					return httpmock.NewStringResponse(400, `{"error","Unexpected format"}`), nil
				}

				subRequest := spec.MemberRequest{}
				if err := json.Unmarshal(bodyBytes, &subRequest); err != nil {
					return httpmock.NewStringResponse(500, `{"error","Failed to unmarshal body to struct"}`), nil
				}

				return httpmock.NewStringResponse(200, MAILCHIMP_MOCK_POST_SUBSCRIBER_200_RESPONSE), nil
			}
			return httpmock.NewStringResponse(401, "{}"), nil
		},
	)
	mc := NewMailchimp("WRONG-us10")

	sub := Subscriber{
		"donal@test.com",
		"Donal",
		"Test",
	}
	err := mc.AddSubscriber(sub, "x1234")
	if err == nil {
		t.Error("Got no error adding subscriber with bad auth token")
	}

	mc = NewMailchimp(MAILCHIMP_TEST_KEY)

	err = mc.AddSubscriber(sub, "x1234")
	if err != nil {
		t.Error("Got error adding subscriber with valid auth token:" + err.Error())
	}

}
