package tools

import (
	"fmt"
	"ggt/types"
	"os"
)

// CheckDir checks if the directory exists.
// If not, it creates it.
func CheckDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}

	return nil
}

func ConfigDirPath() string {
	// get home dir
	dir := os.Getenv("HOME")
	return fmt.Sprintf("%s/%s", dir, types.ConfigDir)
}
func ConfigFilePath() string {
	return fmt.Sprintf("%s/%s", ConfigDirPath(), types.ConfigFile)
}
