package models

import (
	"github.com/byrnedo/tictail_to_mailchimp_forwarder/config"
	"testing"
)

func TestGetLists(t *testing.T) {
	conf := config.Config{}

	if err := conf.ParseFile("../.local.conf"); err != nil {
		t.Error("Failed to read ../local.conf:" + err.Error())
	}

	t.Logf("%v", conf)

	mc := NewMailchimp(&conf)

	data, err := mc.GetLists()
	if err != nil {
		t.Error(err.Error())
	}

	if len(*data) < 1 {
		t.Error("Failed to find lists")
	}
}
