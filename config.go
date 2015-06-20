package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type CiConfg struct {
	Port     string
	Triggers []Trigger
}

type Trigger struct {
	Issuer     string
	Repository string
	Branch     string
	Type       string
	SshUser    string
	SshHost    string
	SshKey     string
	Cmd        string
}

func getConfigContent() ([]byte, error) {
	cfgFilepath := ""

	appPath, err := filepath.Abs(os.Args[0])
	if err != nil {
		return []byte{}, err
	}

	cfgFilepath = appPath + ".json"
	return ioutil.ReadFile(cfgFilepath)
}

func LoadConfig() (*CiConfg, error) {
	cfg := &CiConfg{}

	cfgContent, err := getConfigContent()
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(cfgContent, &cfg)

	return cfg, err
}
