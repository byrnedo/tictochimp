package models

import (
	"github.com/byrnedo/tictochimp/config"
	"github.com/jarcoal/httpmock"
	"net/http"
	"testing"
)

const (
	MAILCHIMP_TEST_URL = "https://api.mailchimp.lol"
	MAILCHIMP_TEST_KEY = "testToken1234"
)

func TestGetLists(t *testing.T) {

	conf := config.Config{}

	if err := conf.ParseFile("../.local.conf"); err != nil {
		t.Error("Failed to read ../local.conf:" + err.Error())
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", MAILCHIMP_TEST_URL,
		func(req *http.Request) (*http.Response, error) {
			if req.Header.Get("Authorization") == "apiKey "+MAILCHIMP_TEST_KEY {
				return httpmock.NewStringResponse(200, MOCK_GET_LISTS_200_RESPONSE), nil
			}
			return httpmock.NewStringResponse(401, "{}"), nil
		},
	)

	t.Logf("%v", conf)

	mc := NewMailchimp(MAILCHIMP_TEST_URL, MAILCHIMP_TEST_KEY)

	data, err := mc.GetLists()
	if err != nil {
		t.Error(err.Error())
	}

	if len(data) != 1 {
		t.Error("Failed to find 1 list")
	}
}

/*
 *
 *func TestAddSubscriberToList(t *testing.T) {
 *
 *    conf := config.Config{}
 *
 *    if err := conf.ParseFile("../.local.conf"); err != nil {
 *        t.Error("Failed to read ../local.conf:" + err.Error())
 *    }
 *
 *    httpmock.Activate()
 *    defer httpmock.DeactivateAndReset()
 *
 *    httpmock.RegisterResponder("GET", conf.Mailchimp.Url+"/3.0/lists/",
 *        httpmock.NewStringResponder(200, MOCK_GET_LISTS_200_RESPONSE))
 *}
 */

const MOCK_GET_LISTS_200_RESPONSE = `{"lists":[{"id":"79aeeb3b51","name":"Nyhetsbrev","contact":{"company":"Helande H\u00e4nder","address1":"H\u00f6torget 2","address2":"","city":"Bor\u00e5s","state":"","zip":"50332","country":"SE","phone":""},"permission_reminder":"","use_archive_bar":true,"campaign_defaults":{"from_name":"Christina Grossi","from_email":"cg@christinagrossi.se","subject":"","language":"en"},"notify_on_subscribe":"","notify_on_unsubscribe":"","date_created":"2015-05-20T09:20:33+00:00","list_rating":3,"email_type_option":false,"subscribe_url_short":"http://eepurl.com/bnZDDX","subscribe_url_long":"http://christinagrossi.us10.list-manage1.com/subscribe?u=4914e9e375e17bb6e666330b1&id=79aeeb3b51","beamer_address":"us10-6137ee8d98-f038eb399a@inbound.mailchimp.com","visibility":"pub","modules":[],"stats":{"member_count":119,"unsubscribe_count":11,"cleaned_count":52,"member_count_since_send":0,"unsubscribe_count_since_send":0,"cleaned_count_since_send":0,"campaign_count":7,"campaign_last_sent":"2015-09-20T19:17:00+00:00","merge_field_count":2,"avg_sub_rate":10,"avg_unsub_rate":3,"target_sub_rate":20,"open_rate":55.711422845691,"click_rate":13.452914798206,"last_sub_date":"2015-09-20T07:54:25+00:00","last_unsub_date":"2015-09-12T16:45:12+00:00"},"_links":[{"rel":"self","href":"https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51","method":"GET","targetSchema":"https://us10.api.mailchimp.com/schema/3.0/Lists/Instance.json"},{"rel":"parent","href":"https://us10.api.mailchimp.com/3.0/lists","method":"GET","targetSchema":"https://us10.api.mailchimp.com/schema/3.0/Lists/Collection.json","schema":"https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists.json"},{"rel":"update","href":"https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51","method":"PATCH","schema":"https://us10.api.mailchimp.com/schema/3.0/Lists/Instance.json"},{"rel":"delete","href":"https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51","method":"DELETE"},{"rel":"abuse-reports","href":"https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/abuse-reports","method":"GET","targetSchema":"https://us10.api.mailchimp.com/schema/3.0/Lists/Abuse/Collection.json","schema":"https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists/Abuse.json"},{"rel":"activity","href":"https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/activity","method":"GET","targetSchema":"https://us10.api.mailchimp.com/schema/3.0/Lists/Activity/Collection.json"},{"rel":"clients","href":"https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/clients","method":"GET","targetSchema":"https://us10.api.mailchimp.com/schema/3.0/Lists/Clients/Collection.json"},{"rel":"growth-history","href":"https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/growth-history","method":"GET","targetSchema":"https://us10.api.mailchimp.com/schema/3.0/Lists/Growth/Collection.json","schema":"https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists/Growth.json"},{"rel":"interest-categories","href":"https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/interest-categories","method":"GET","targetSchema":"https://us10.api.mailchimp.com/schema/3.0/Lists/InterestCategories/Collection.json","schema":"https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists/InterestCategories.json"},{"rel":"members","href":"https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/members","method":"GET","targetSchema":"https://us10.api.mailchimp.com/schema/3.0/Lists/Members/Collection.json","schema":"https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists/Members.json"},{"rel":"merge-fields","href":"https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/merge-fields","method":"GET","targetSchema":"https://us10.api.mailchimp.com/schema/3.0/Lists/MergeFields/Collection.json","schema":"https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists/MergeFields.json"},{"rel":"segments","href":"https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/segments","method":"GET","targetSchema":"https://us10.api.mailchimp.com/schema/3.0/Lists/Segments/Collection.json","schema":"https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists/Segments.json"}]}]}`
