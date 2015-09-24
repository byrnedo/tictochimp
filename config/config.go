package config

import (
	"errors"
	"github.com/liyinhgqw/typesafe-config/parse"
	"io/ioutil"
	"os"
	"regexp"
)

type Config struct {
	Tictail struct {
		Url          string
		ClientID     string
		ClientSecret string
		Product      string
	}
	Mailchimp struct {
		Url         string
		AccessToken string
		List        string
	}
	underlyingData *parse.Tree
}

func (c *Config) ParseFile(path string) error {
	os.Stdout.WriteString("Parsing config file\n")
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.New("Failed to read config file")
	}
	return c.Parse(bytes)
}

func (c *Config) Parse(configFileData []byte) (err error) {
	os.Stdout.WriteString("Parsing config string\n")
	c.underlyingData, err = parse.Parse("config", string(configFileData))
	if err == nil {
		populateConfigVars(c)
	}
	return
}

func populateConfigVars(c *Config) {
	os.Stdout.WriteString("Parsing config vars\n")
	typesafeConf := c.underlyingData.GetConfig()

	setDefaultValues(c)

	if val, err := typesafeConf.GetString("tictail.url"); err == nil {
		c.Tictail.Url = unquoteString(val)
	}

	if val, err := typesafeConf.GetString("tictail.client-id"); err == nil {
		c.Tictail.ClientID = unquoteString(val)
	}

	if val, err := typesafeConf.GetString("tictail.client-secret"); err == nil {
		c.Tictail.ClientSecret = unquoteString(val)
	}

	if val, err := typesafeConf.GetString("tictail.product"); err == nil {
		c.Tictail.Product = unquoteString(val)
	}

	if val, err := typesafeConf.GetString("mailchimp.url"); err == nil {
		c.Mailchimp.Url = unquoteString(val)
	}

	if val, err := typesafeConf.GetString("mailchimp.access-token"); err == nil {
		c.Mailchimp.AccessToken = unquoteString(val)
	}

	if val, err := typesafeConf.GetString("mailchimp.list"); err == nil {
		c.Mailchimp.List = unquoteString(val)
	}
}

func unquoteString(value string) string {
	re := regexp.MustCompile("^\"(.*)\"$")
	if strippedVal := re.FindStringSubmatch(value); strippedVal != nil {
		return strippedVal[1]
	} else {
		return value
	}
}

func setDefaultValues(c *Config) {
	c.Tictail.Url = "https://tictail.com"
	c.Mailchimp.Url = "https://us10.api.mailchimp.com/3.0"
}

func (c *Config) GetUnderlyingData() (data *parse.Tree) {
	return c.underlyingData
}
