# Tictochimp

Adds the email address to a mailchimp list from an order for a specific product.

For example, if you wished to deliver a mail campaign based on a purchase through tictail.

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
