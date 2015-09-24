package main

import (
	"flag"
	"github.com/byrnedo/tictochimp/config"
	"io/ioutil"
	"os"
)

var (
	configFile string
	showUsage  bool
)

func main() {

	flag.StringVar(&configFile, "conf", "", "Configuration file path")
	flag.BoolVar(&showUsage, "help", false, "Show usage information")
	flag.Parse()

	if len(configFile) == 0 {
		showUsage = true
	}

	if showUsage {
		flag.Usage()
		os.Exit(1)
	}

	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		os.Stderr.WriteString("Error opening config file:" + err.Error())
		os.Exit(1)
	}

	var configData = config.Config{}

	err = configData.Parse(bytes)
	if err != nil {
		os.Stderr.WriteString("Error parsing config file:" + err.Error())
		os.Exit(1)
	} else {
		os.Stdout.WriteString("Got config:" + configData.GetUnderlyingData().Root.String())
		//conf := result.GetConfig()
	}
}
