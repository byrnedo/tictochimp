package config

import (
	"testing"
)

func TestParse(t *testing.T) {
	conf := Config{}

	t.Logf("Conf object before: %v", conf)

	if err := conf.ParseFile("../test.conf"); err != nil {
		t.Error("Failed to read ../test.conf:" + err.Error())
	}

	if conf.Tictail.ClientID != "test" {
		t.Error("Incorrect ClientID value")
	}

	t.Log(conf.Mailchimp.Url)

	if conf.Mailchimp.Url != "https://us10.api.mailchimp.com" {
		t.Error("Incorrect Url value")
	}

	t.Logf("Conf object after: %v", conf)

}
