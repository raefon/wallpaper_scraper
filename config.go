package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	APIKey        string `json:"api_key"`
	SearchQuery   string `json:"search_query"`
	Resolution    string `json:"resolution"`
	Categories    string `json:"categories"`
	Purity        string `json:"purity"`
	Sorting       string `json:"sorting"`
	Order         string `json:"order"`
	DownloadLimit int    `json:"download_limit"`
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
