package helper

import (
	"encoding/json"
	"ggt/config"
	"ggt/types"
	"log"
	"os"
)

// GetConfig get config from config file
func GetConfig() (config.Cfg, error) {
	var cf config.Cfg
	// read config
	_, err := os.Stat(types.ConfigPath)
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(types.NeedLoginStatus)
		return cf, err
	}

	// read string from f
	// unmarshal string to cf
	buf, err := os.ReadFile(types.ConfigPath)
	if err != nil {
		log.Fatal(err)
		return cf, err
	}

	err = json.Unmarshal(buf, &cf)
	if err != nil {
		log.Fatal(err)
		return cf, err
	}

	return cf, nil
}
