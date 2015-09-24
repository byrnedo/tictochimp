package config

import (
	"errors"
	"github.com/liyinhgqw/typesafe-config/parse"
	"io/ioutil"
)

type Config struct {
	Tictail struct {
		Url          string
		ClientKey    string
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
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.New("Failed to read config file")
	}
	return c.Parse(bytes)
}

func (c *Config) Parse(configFileData []byte) (err error) {
	c.underlyingData, err = parse.Parse("config", string(configFileData))

	if err != nil {
		populateConfigVars(c)
	}
	return
}

func setDefaultValues(c *Config) {
	c.Tictail.Url = "https://tictail.com"
	c.Mailchimp.Url = "https://us10.api.mailchimp.com/3.0"
}

func populateConfigVars(c *Config) {
	typesafeConf := c.underlyingData.GetConfig()

	setDefaultValues(c)

	if val, err := typesafeConf.GetString("tictail.url"); err != nil {
		c.Tictail.Url = val
	}

	if val, err := typesafeConf.GetString("tictail.client-key"); err != nil {
		c.Tictail.ClientKey = val
	}

	if val, err := typesafeConf.GetString("tictail.client-secret"); err != nil {
		c.Tictail.ClientSecret = val
	}

	if val, err := typesafeConf.GetString("tictail.product"); err != nil {
		c.Tictail.Product = val
	}

	if val, err := typesafeConf.GetString("mailchimp.url"); err != nil {
		c.Mailchimp.Url = val
	}

	if val, err := typesafeConf.GetString("mailchimp.access-token"); err != nil {
		c.Mailchimp.AccessToken = val
	}

	if val, err := typesafeConf.GetString("mailchimp.list"); err != nil {
		c.Mailchimp.List = val
	}
}

func (c *Config) GetUnderlyingData() (data *parse.Tree) {
	return c.underlyingData
}
