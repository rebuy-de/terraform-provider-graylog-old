package testutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func AssertGoldenJSON(filename string, data interface{}) error {
	generated, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	generated = append(generated, '\n')

	return AssertGolden(filename, generated)
}

func AssertGolden(filename string, data []byte) error {
	if v := os.Getenv("TF_UPDATE_GOLDEN"); v == "1" {
		err := ioutil.WriteFile(filename, data, os.FileMode(0644))
		if err != nil {
			return err
		}
	}

	golden, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	if string(golden) != string(data) {
		return fmt.Errorf("Generated file '%s' doesn't match golden file. Update by setting 'TF_UPDATE_GOLDEN=1'.", filename)
	}

	return nil
}
