package mailchimp

const MAILCHIMP_MOCK_POST_SUBSCRIBER_200_RESPONSE = `{"status": "success"}`

const MAILCHIMP_MOCK_GET_LISTS_200_RESPONSE = `{
  "lists": [
    {
      "id": "xxxtestid",
      "name": "Mirakelkort",
      "contact": {
        "company": "TestCompany",
        "address1": "Test 2",
        "address2": "",
        "city": "",
        "state": "",
        "zip": "",
        "country": "SE",
        "phone": ""
      },
      "permission_reminder": "",
      "use_archive_bar": true,
      "campaign_defaults": {
        "from_name": "Test User",
        "from_email": "test@testuser.com",
        "subject": "",
        "language": "en"
      },
      "notify_on_subscribe": "",
      "notify_on_unsubscribe": "",
      "date_created": "2015-09-23T11:07:23+00:00",
      "list_rating": 0,
      "email_type_option": false,
      "subscribe_url_short": "http://eepurl.com/xxx",
      "subscribe_url_long": "http://testuser.us10.list-manage.com/subscribe?u=xxx&id=xxxtestid",
      "beamer_address": "",
      "visibility": "pub",
      "modules": [],
      "stats": {
        "member_count": 1,
        "unsubscribe_count": 0,
        "cleaned_count": 0,
        "member_count_since_send": 1,
        "unsubscribe_count_since_send": 0,
        "cleaned_count_since_send": 0,
        "campaign_count": 34,
        "campaign_last_sent": "",
        "merge_field_count": 2,
        "avg_sub_rate": 0,
        "avg_unsub_rate": 0,
        "target_sub_rate": 0,
        "open_rate": 0,
        "click_rate": 0,
        "last_sub_date": "2015-09-24T18:36:10+00:00",
        "last_unsub_date": ""
      },
      "_links": [
        {
          "rel": "self",
          "href": "https://us10.api.mailchimp.com/3.0/lists/xxxListId",
          "method": "GET",
          "targetSchema": "https://us10.api.mailchimp.com/schema/3.0/Lists/Instance.json"
        },
        {
          "rel": "parent",
          "href": "https://us10.api.mailchimp.com/3.0/lists",
          "method": "GET",
          "targetSchema": "https://us10.api.mailchimp.com/schema/3.0/Lists/Collection.json",
          "schema": "https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists.json"
        },
        {
          "rel": "update",
          "href": "https://us10.api.mailchimp.com/3.0/lists/xxxListId",
          "method": "PATCH",
          "schema": "https://us10.api.mailchimp.com/schema/3.0/Lists/Instance.json"
        },
        {
          "rel": "delete",
          "href": "https://us10.api.mailchimp.com/3.0/lists/xxxListId",
          "method": "DELETE"
        },
        {
          "rel": "abuse-reports",
          "href": "https://us10.api.mailchimp.com/3.0/lists/xxxListId/abuse-reports",
          "method": "GET",
          "targetSchema": "https://us10.api.mailchimp.com/schema/3.0/Lists/Abuse/Collection.json",
          "schema": "https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists/Abuse.json"
        },
        {
          "rel": "activity",
          "href": "https://us10.api.mailchimp.com/3.0/lists/xxxListId/activity",
          "method": "GET",
          "targetSchema": "https://us10.api.mailchimp.com/schema/3.0/Lists/Activity/Collection.json"
        },
        {
          "rel": "clients",
          "href": "https://us10.api.mailchimp.com/3.0/lists/xxxListId/clients",
          "method": "GET",
          "targetSchema": "https://us10.api.mailchimp.com/schema/3.0/Lists/Clients/Collection.json"
        },
        {
          "rel": "growth-history",
          "href": "https://us10.api.mailchimp.com/3.0/lists/xxxListId/growth-history",
          "method": "GET",
          "targetSchema": "https://us10.api.mailchimp.com/schema/3.0/Lists/Growth/Collection.json",
          "schema": "https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists/Growth.json"
        },
        {
          "rel": "interest-categories",
          "href": "https://us10.api.mailchimp.com/3.0/lists/xxxListId/interest-categories",
          "method": "GET",
          "targetSchema": "https://us10.api.mailchimp.com/schema/3.0/Lists/InterestCategories/Collection.json",
          "schema": "https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists/InterestCategories.json"
        },
        {
          "rel": "members",
          "href": "https://us10.api.mailchimp.com/3.0/lists/xxxListId/members",
          "method": "GET",
          "targetSchema": "https://us10.api.mailchimp.com/schema/3.0/Lists/Members/Collection.json",
          "schema": "https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists/Members.json"
        },
        {
          "rel": "merge-fields",
          "href": "https://us10.api.mailchimp.com/3.0/lists/xxxListId/merge-fields",
          "method": "GET",
          "targetSchema": "https://us10.api.mailchimp.com/schema/3.0/Lists/MergeFields/Collection.json",
          "schema": "https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists/MergeFields.json"
        },
        {
          "rel": "segments",
          "href": "https://us10.api.mailchimp.com/3.0/lists/xxxListId/segments",
          "method": "GET",
          "targetSchema": "https://us10.api.mailchimp.com/schema/3.0/Lists/Segments/Collection.json",
          "schema": "https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists/Segments.json"
        }
      ]
    }
  ],
  "_links": [
    {
      "rel": "self",
      "href": "https://us10.api.mailchimp.com/3.0/lists",
      "method": "GET",
      "targetSchema": "https://us10.api.mailchimp.com/schema/3.0/Lists/Collection.json",
      "schema": "https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists.json"
    },
    {
      "rel": "parent",
      "href": "https://us10.api.mailchimp.com/3.0/",
      "method": "GET",
      "targetSchema": "https://us10.api.mailchimp.com/schema/3.0/Root.json"
    },
    {
      "rel": "create",
      "href": "https://us10.api.mailchimp.com/3.0/lists",
      "method": "POST",
      "schema": "https://us10.api.mailchimp.com/schema/3.0/Lists/Instance.json"
    }
  ],
  "total_items": 1
}`

const MAILCHIMP_MOCK_GET_LIST_MEMBERS_200_RESPONSE = `{
"members" : [{
      "id": "bbcf38a9cd250024dc08730de8e606d5",
      "email_address": "testuser@gmail.com",
      "unique_email_id": "bb6c90d60a",
      "email_type": "html",
      "status": "subscribed",
      "merge_fields": {
        "FNAME": "Test",
        "LNAME": "User"
      },
      "stats": {
        "avg_open_rate": 1,
        "avg_click_rate": 0
      },
      "ip_signup": "",
      "timestamp_signup": "",
      "ip_opt": "77.218.231.142",
      "timestamp_opt": "2015-08-25 19:08:48",
      "member_rating": 4,
      "last_changed": "2015-08-25 19:08:48",
      "language": "",
      "vip": false,
      "email_client": "",
      "location": {
        "latitude": 59.3114,
        "longitude": 18.0862,
        "gmtoff": 1,
        "dstoff": 2,
        "country_code": "SE",
        "timezone": "Europe/Prague"
      },
      "list_id": "79aeeb3b51",
      "_links": [
        {
          "rel": "self",
          "href": "https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/members/bbcf38a9cd250024dc08730de8e606d5",
          "method": "GET",
          "targetSchema": "https://us10.api.mailchimp.com/schema/3.0/Lists/Members/Instance.json"
        },
        {
          "rel": "parent",
          "href": "https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/members",
          "method": "GET",
          "targetSchema": "https://us10.api.mailchimp.com/schema/3.0/Lists/Members/Collection.json",
          "schema": "https://us10.api.mailchimp.com/schema/3.0/CollectionLinks/Lists/Members.json"
        },
        {
          "rel": "update",
          "href": "https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/members/bbcf38a9cd250024dc08730de8e606d5",
          "method": "PATCH",
          "schema": "https://us10.api.mailchimp.com/schema/3.0/Lists/Members/Instance.json"
        },
        {
          "rel": "upsert",
          "href": "https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/members/bbcf38a9cd250024dc08730de8e606d5",
          "method": "PUT",
          "schema": "https://us10.api.mailchimp.com/schema/3.0/Lists/Members/Instance.json"
        },
        {
          "rel": "delete",
          "href": "https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/members/bbcf38a9cd250024dc08730de8e606d5",
          "method": "DELETE"
        },
        {
          "rel": "activity",
          "href": "https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/members/bbcf38a9cd250024dc08730de8e606d5/activity",
          "method": "GET",
          "targetSchema": "https://us10.api.mailchimp.com/schema/3.0/Lists/Members/Activity/Collection.json"
        },
        {
          "rel": "goals",
          "href": "https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/members/bbcf38a9cd250024dc08730de8e606d5/goals",
          "method": "GET",
          "targetSchema": "https://us10.api.mailchimp.com/schema/3.0/Lists/Members/Goals/Collection.json"
        },
        {
          "rel": "notes",
          "href": "https://us10.api.mailchimp.com/3.0/lists/79aeeb3b51/members/bbcf38a9cd250024dc08730de8e606d5/notes",
          "method": "GET",
          "targetSchema": "https://us10.api.mailchimp.com/schema/3.0/Lists/Members/Notes/Collection.json"
        }
      ]
    }],
	"total_items" : 1
}`
