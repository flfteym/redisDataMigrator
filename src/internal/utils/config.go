package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func InitConfig(configFile string) Config {
	rdmConfig := Config{}
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, &rdmConfig)
	if err != nil {
		panic(err)
	}

	return rdmConfig
}
