package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type Config struct {
	Inputs  []Pin `json:"inputs"`
	Outputs []Pin `json:"outputs"`
	Port    int   `json:"port"`
}

func NewConfigDefault() Config {
	return Config{
		Inputs: []Pin{
			{Name: "GPIO1", Register: 1001},
			{Name: "GPIO2", Register: 1002},
			{Name: "GPIO3", Register: 1003},
			{Name: "GPIO4", Register: 1004},
			{Name: "GPIO5", Register: 1005},
			{Name: "GPIO6", Register: 1006},
			{Name: "GPIO7", Register: 1007},
			{Name: "GPIO8", Register: 1008},
			{Name: "GPIO9", Register: 1009},
			{Name: "GPIO10", Register: 1010},
			{Name: "GPIO11", Register: 1011},
			{Name: "GPIO12", Register: 1012},
			{Name: "GPIO13", Register: 1013},
		},
		Outputs: []Pin{
			{Name: "GPIO14", Register: 2014},
			{Name: "GPIO15", Register: 2015},
			{Name: "GPIO16", Register: 2016},
			{Name: "GPIO17", Register: 2017},
			{Name: "GPIO18", Register: 2018},
			{Name: "GPIO19", Register: 2019},
			{Name: "GPIO20", Register: 2020},
			{Name: "GPIO21", Register: 2021},
			{Name: "GPIO22", Register: 2022},
			{Name: "GPIO23", Register: 2023},
			{Name: "GPIO24", Register: 2024},
			{Name: "GPIO25", Register: 2025},
			{Name: "GPIO26", Register: 2026},
			{Name: "GPIO27", Register: 2027},
		},
		Port: 1502,
	}
}

var configFile = "config.json"

func ReadConfig() (conf Config, err error) {
	if _, e := os.Stat(configFile); errors.Is(e, os.ErrNotExist) {
		conf = NewConfigDefault()
		file, _ := json.MarshalIndent(conf, "", "    ")
		err = os.WriteFile(configFile, file, 0644)
		if err == nil {
			return conf, fmt.Errorf("config file does not exist")
		}
		return
	}

	jsonFile, err := os.Open(configFile)
	defer jsonFile.Close()

	if err != nil {
		return
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return
	}
	err = json.Unmarshal(byteValue, &conf)

	if len(conf.Inputs) <= 0 {
		return conf, fmt.Errorf("config file for inputs is empty")
	}

	if len(conf.Outputs) <= 0 {
		return conf, fmt.Errorf("config file for outputs is empty")
	}

	if conf.Port < 0 || conf.Port > 65535 {
		return conf, fmt.Errorf("port out of range")
	}
	return
}
