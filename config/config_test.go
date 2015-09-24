package config

import (
	"testing"
)

func TestParse(t *testing.T) {
	conf := config.Config{}

	if err := conf.ParseFile("../.local.conf"); err != nil {
		t.Error("Failed to read ../local.conf:" + err.Error())
	}

	t.Logf("%v", conf)

}
