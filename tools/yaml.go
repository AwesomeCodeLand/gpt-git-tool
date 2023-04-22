package tools

import (
	"ggt/config"

	"gopkg.in/yaml.v2"
)

func Config(data []byte) (cf config.Cfg, err error) {
	err = yaml.Unmarshal(data, &cf)
	if err != nil {
		return cf, err
	}
	return cf, nil
}

func Marshal(cf config.Cfg) (data []byte, err error) {
	data, err = yaml.Marshal(cf)
	if err != nil {
		return data, err
	}
	return data, nil
}
