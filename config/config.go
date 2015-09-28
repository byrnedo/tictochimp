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
		AccessToken string
		ProductName string
		StoreName   string
	}
	Mailchimp struct {
		AccessToken string
		ListName    string
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

func populateConfigVars(c *Config) error {
	os.Stdout.WriteString("Parsing config vars\n")
	typesafeConf := c.underlyingData.GetConfig()

	setDefaultValues(c)

	if val, err := typesafeConf.GetString("tictail.access-token"); err == nil {
		c.Tictail.AccessToken = unquoteString(val)
		if c.Tictail.AccessToken == "" {
			return errors.New("tictail.access-token: can not be empty")
		}
	}

	if val, err := typesafeConf.GetString("tictail.store-name"); err == nil {
		c.Tictail.StoreName = unquoteString(val)
		if c.Tictail.StoreName == "" {
			return errors.New("tictail.store-name: can not be empty")
		}
	}

	if val, err := typesafeConf.GetString("tictail.product-name"); err == nil {
		c.Tictail.ProductName = unquoteString(val)
		if c.Tictail.ProductName == "" {
			return errors.New("tictail.product-name: can not be empty")
		}
	}

	if val, err := typesafeConf.GetString("mailchimp.access-token"); err == nil {
		c.Mailchimp.AccessToken = unquoteString(val)
		if c.Mailchimp.AccessToken == "" {
			return errors.New("mailchimp.access-token: can not be empty")
		}
	}

	if val, err := typesafeConf.GetString("mailchimp.list-name"); err == nil {
		c.Mailchimp.ListName = unquoteString(val)
		if c.Mailchimp.ListName == "" {
			return errors.New("mailchimp.list-name: can not be empty")
		}
	}
	return nil
}

func validateNotEmpty(value string) {
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
}

func (c *Config) GetUnderlyingData() (data *parse.Tree) {
	return c.underlyingData
}
