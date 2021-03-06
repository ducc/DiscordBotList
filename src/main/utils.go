package main

import (
	"encoding/json"
	"io/ioutil"
)

func loadConfig(path string) (*config, error) {
	var conf config
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
