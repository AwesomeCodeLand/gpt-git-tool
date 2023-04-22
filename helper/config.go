package helper

import (
	"ggt/config"
	"ggt/tools"
	"ggt/types"
	"log"
	"os"
)

// GetConfig get config from config file
func GetConfig() (config.Cfg, error) {
	var cf config.Cfg
	// read config
	_, err := os.Stat(tools.ConfigFilePath())
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(types.NeedLoginStatus)
		return cf, err
	}

	// read string from f
	// unmarshal string to cf
	buf, err := os.ReadFile(tools.ConfigFilePath())
	if err != nil {
		log.Fatal(err)
		return cf, err
	}

	return tools.Config(buf)
}

// SaveConfig save config to config file
func SaveConfig(cf config.Cfg) error {

	err := tools.CheckDir(tools.ConfigDirPath())
	if err != nil {
		return err
	}

	// marshal cf to string
	buf, err := tools.Marshal(cf)
	if err != nil {
		return err
	}

	os.Remove(tools.ConfigFilePath())
	// write string to f, if f not exist, create it
	f, err := os.OpenFile(tools.ConfigFilePath(), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(buf)
	// err = os.WriteFile(types.ConfigPath, buf, 0644)
	if err != nil {
		return err
	}

	return nil
}
