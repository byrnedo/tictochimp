#!/bin/bash

source .local.source

curl -s -m 15 --header "Authorization: apiKey ${mailchimp_api_key}"  https://${mailchimp_datacenter}.api.mailchimp.com/3.0/lists/${mailchimp_list_id}/members
