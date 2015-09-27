# Tictochimp

Adds the email address from a Tictail purchase to a Mailchimp mailing list.

For example, if you wished to deliver a mail campaign based on a purchase of specific product through Tictail.

## To build

You must have go installed.

If you do then just build normally:

    cd tictochimp/
    go get
    go build 

There are only 2 external dependencies ( 1 for testing ) so should be quick. 

## Usage

    ./tictochimp -conf <config file name>

    ## Or to do a dry run

    ./tictochimp -conf <config file name> -dry-run

## Config file format

Config file is HOCON based, only a few fields in it.
Here's an example config:

    tictail {
        access-token = "x1234-us10"
        product-name = "Some product title"
        store-name = "yourstoresubdomain"
    }

    mailchimp {
        access-token = "x1234"
        list-name = "mymaillist"
    }

NOTE: tictail access tokens must have the '-us10' or similar suffix included in order to determine the api url.
