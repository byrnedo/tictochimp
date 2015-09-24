package models

import (
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

const (
	MAILCHIMP_TEST_URL = "https://api.mailchimp.lol"
	MAILCHIMP_TEST_KEY = "testToken1234"
)

func TestGetLists(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", MAILCHIMP_TEST_URL+"/3.0/lists",
		func(req *http.Request) (*http.Response, error) {
			if req.Header.Get("Authorization") == "apiKey "+MAILCHIMP_TEST_KEY {
				return httpmock.NewStringResponse(200, MOCK_GET_LISTS_200_RESPONSE), nil
			}
			return httpmock.NewStringResponse(401, "{}"), nil
		},
	)

	mc := NewMailchimp(MAILCHIMP_TEST_URL, MAILCHIMP_TEST_KEY)

	data, err := mc.GetLists()
	if err != nil {
		t.Error(err.Error())
	}

	if len(data) != 1 {
		t.Error("Failed to find 1 list")
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

				subRequest := memberRequest{}
				if err := json.Unmarshal(bodyBytes, &subRequest); err != nil {
					return httpmock.NewStringResponse(500, `{"error","Failed to unmarshal body to struct"}`), nil
				}

				return httpmock.NewStringResponse(200, MOCK_POST_SUBSCRIBER_200_RESPONSE), nil
			}
			return httpmock.NewStringResponse(401, "{}"), nil
		},
	)
	mc := NewMailchimp(MAILCHIMP_TEST_URL, "WRONG")

	sub := Subscriber{
		"donal@test.com",
		"Donal",
		"Test",
	}
	err := mc.AddSubscriber(sub, "x1234")
	if err == nil {
		t.Error("Got no error adding subscriber with bad auth token")
	}

	mc = NewMailchimp(MAILCHIMP_TEST_URL, MAILCHIMP_TEST_KEY)

	err = mc.AddSubscriber(sub, "x1234")
	if err != nil {
		t.Error("Got error adding subscriber with valid auth token:" + err.Error())
	}

}

const MOCK_POST_SUBSCRIBER_200_RESPONSE = `{"status": "success"}`

const MOCK_GET_LISTS_200_RESPONSE = `{"lists":[{"id":"79aeeb3b51","name":"Nyhetsbrev","contact":{"company":"Helande H\u00e4nder","address1":"H\u00f6torget 2","address2":"","city":"Bor\u00e5s","state":"","zip":"50332","country":"SE","phone":""},"permission_reminder":"","use_archive_bar":true,"campaign_defaults":{"from_name":"Christina Grossi","from_email":"cg@christinagrossi.se","subject":"","language":"en"},"notify_on_subscribe":"","notify_on_unsubscribe":"","date_created":"2015-05-20T09:20:33+00:00","list_rating":3,"email_type_option":false,"subscribe_url_short":"http://eepurl.com/bnZDDX","subscribe_url_long":"http://christinagrossi.us10.list-manage1.com/subscribe?u=4914e9e375e17bb6e666330b1&id=79aeeb3b51","beamer_address":"us10-6137ee8d98-f038eb399a@inbound.mailchimp.com","visibility":"pub","modules":[],"stats":{"member_count":119,"unsubscribe_count":11,"cleaned_count":52,"member_count_since_send":0,"unsubscribe_count_since_send":0,"cleaned_count_since_send":0,"campaign_count":7,"campaign_last_sent":"2015-09-20T19:17:00+00:00","merge_field_count":2,"avg_sub_rate":10,"avg_unsub_rate":3,"target_sub_rate":20,"open_rate":55.711422845691,"click_rate":13.452914798206,"last_sub_date":"2015-09-20T07:54:25+00:00","last_unsub_date":"2015-09-12T16:45:12+00:00"},"_links":[{"rel":"self","href":"https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51","method":"GET","targetSchema":"https://us10.api.mailchimp.com/schema/3.0/Lists/Instance.json"},{"rel":"parent","href":"https://us10.api.mailchimp.com/3.0/lists","method":"GET","targetSchema":"https://us10.api.mailchimp.com/schema/3.0/Lists/Collection.json","schema":"https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists.json"},{"rel":"update","href":"https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51","method":"PATCH","schema":"https://us10.api.mailchimp.com/schema/3.0/Lists/Instance.json"},{"rel":"delete","href":"https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51","method":"DELETE"},{"rel":"abuse-reports","href":"https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/abuse-reports","method":"GET","targetSchema":"https://us10.api.mailchimp.com/schema/3.0/Lists/Abuse/Collection.json","schema":"https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists/Abuse.json"},{"rel":"activity","href":"https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/activity","method":"GET","targetSchema":"https://us10.api.mailchimp.com/schema/3.0/Lists/Activity/Collection.json"},{"rel":"clients","href":"https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/clients","method":"GET","targetSchema":"https://us10.api.mailchimp.com/schema/3.0/Lists/Clients/Collection.json"},{"rel":"growth-history","href":"https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/growth-history","method":"GET","targetSchema":"https://us10.api.mailchimp.com/schema/3.0/Lists/Growth/Collection.json","schema":"https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists/Growth.json"},{"rel":"interest-categories","href":"https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/interest-categories","method":"GET","targetSchema":"https://us10.api.mailchimp.com/schema/3.0/Lists/InterestCategories/Collection.json","schema":"https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists/InterestCategories.json"},{"rel":"members","href":"https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/members","method":"GET","targetSchema":"https://us10.api.mailchimp.com/schema/3.0/Lists/Members/Collection.json","schema":"https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists/Members.json"},{"rel":"merge-fields","href":"https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/merge-fields","method":"GET","targetSchema":"https://us10.api.mailchimp.com/schema/3.0/Lists/MergeFields/Collection.json","schema":"https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists/MergeFields.json"},{"rel":"segments","href":"https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/segments","method":"GET","targetSchema":"https://us10.api.mailchimp.com/schema/3.0/Lists/Segments/Collection.json","schema":"https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists/Segments.json"}]}]}`
