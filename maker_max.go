package main

import (
	"os"
	"encoding/json"

	"io/ioutil"
	"path/filepath"
)

func MakeDir(name string) error {
	absPath, err := filepath.Abs(name)
	if err != nil {
		return err
	}

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		err := os.Mkdir(name, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func ParseSeed(path string) (*ChimpCodeSeed, error) {
	seed := new(ChimpCodeSeed)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return seed, err
	}
	err = json.Unmarshal(data, seed)
	return seed, err
}
