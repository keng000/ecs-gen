package json

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// AppendWrite dumps a data into a json file or AppendWrite if already exists
func AppendWrite(path string, data map[string]string) error {
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

// Read is
func Read(path string) (map[string]string, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var dump map[string]string
	if err := json.Unmarshal(contents, &dump); err != nil {
		return nil, err
	}

	return dump, nil
}

// StructToMap is
func StructToMap(i interface{}) (map[string]string, error) {

	return nil, nil
}
