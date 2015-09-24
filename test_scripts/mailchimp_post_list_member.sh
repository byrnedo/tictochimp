#!/bin/bash

source .local.source

fname=John
lname=Smith
email_address=john@smith.com
request="{ \"email_address\": \"${email_address}\", \"email_type\": \"html\", \"status_if_new\": \"subscribed\", \"status\": \"subscribed\", \"merge_fields\": { \"FNAME\": \"${fname}\", \"LNAME\": \"${lname}\"}, \"language\":\"sv\", \"vip\": false }"
echo $request

curl -s -m 15 -X POST --header "Authorization: apiKey ${api_key}"  --data "${request}" https://${datacenter}.api.mailchimp.com/3.0/lists/${mailchimp_list_id}/members


#{
#    "email_type": "",
#    "status": "",
#    "status_if_new": "",
#    "merge_fields": {},
#    "interests": {},
#    "stats": {},
#    "language": "",
#    "vip": "",
#    "location": {
#        "latitude": "",
#        "longitude": ""
#    }
#}
# Example
# {
#      "email_address": "magnus76@live.se",
#      "email_type": "html",
#      "status_if_new": "subscribed",
#      "status": "subscribed",
#      "merge_fields": {
#        "FNAME": "Magnus",
#        "LNAME": ""
#      },
#      "language": "sv",
#      "vip": false
# }
#


