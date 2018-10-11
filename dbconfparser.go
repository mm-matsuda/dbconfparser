package dbconfparser

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Parse dbconf.yml to string
func Parse(path string, env string) (string, string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", "", err
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(bytes, &m)
	if err != nil {
		return "", "", err
	}

	driver := m[env].(map[interface{}]interface{})["driver"].(string)
	open := m[env].(map[interface{}]interface{})["open"].(string)

	return driver, open, nil
}
