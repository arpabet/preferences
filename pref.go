package preferences

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	groupName = "Arpabet"
)

func MakeAppDataDir(appName string) (string, error) {
	dir := AppDataDir(appName)
	err := os.MkdirAll(dir, 0700)
	return dir, err
}

func LoadJsonFile(appName, fileName string, v interface{}) (bool, error) {
	dir := AppDataDir(appName)
	fullPath := filepath.Join(dir, fileName)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return false, nil
	}
	blob, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal(blob, v)
	return err == nil, err
}

func SaveJsonFile(appName, fileName string, v interface{}) error {
	blob, err := json.Marshal(v)
	if err != nil {
		return err
	}
	dir, err := MakeAppDataDir(appName)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join(dir, fileName), blob, 600)
}