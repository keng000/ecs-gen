package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Append dumps a data into a json file or append if already exists
func Append(path string, data map[string]interface{}) error {
	if filepath.Ext(path) != ".json" {
		return errors.New("path extension should be `.json`")
	}

	_, err := os.Stat(path)
	if err == nil {
		// Exist
		dump, err := Read(path)
		if err != nil {
			return err
		}

		for key, val := range dump {
			data[key] = val
		}
	}

	fp, err := os.Create(path)
	if err != nil {
		return err
	}

	outputContents, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	_, err = fp.Write(outputContents)
	if err != nil {
		return err
	}

	return nil
}

func Read(path string) (map[string]interface{}, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var dump map[string]interface{}
	if err := json.Unmarshal(contents, &dump); err != nil {
		return nil, err
	}

	return dump, nil
}
