#!/bin/bash

source .local.source

cmd="curl -s -m 10 'https://tictail.com/oauth/token' \
    -d \"client_id=${tictail_client_id}\" \
    -d \"client_secret=${tictail_client_secret}\" \
    -d 'grant_type=password' \
    -d \"username=${tictail_user}\" \
    -d \"password=${tictail_pass}\""

echo $cmd
