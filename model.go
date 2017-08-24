package main

import (
	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Key string `json:"key"`
	Cx  string `json:"cx"`
}

func (cfg *Config) WriteFile(path string) error {
	dir := filepath.Dir(path)
	if err := mkdir(dir); err != nil {
		return errors.Wrap(err, "unable to create config directory")
	}

	out, err := yaml.Marshal(cfg)
	if err != nil {
		return errors.Wrap(err, "unable to marshal as yaml file")
	}

	if err := ioutil.WriteFile(path, out, 0655); err != nil {
		return errors.Wrap(err, "unable to write data to file")
	}

	return nil
}

func GetConf(path string) (*Config, error) {
	cfg := &Config{}

	in, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read config file")
	}

	if err := yaml.Unmarshal(in, &cfg); err != nil {
		return nil, errors.Wrap(err, "unable to parse config file")
	}

	return cfg, nil
}

type Result struct {
	Queries     *Query    `json:"queries"`
	Info        *Info     `json:"searchInformation"`
	Corrections *Spelling `json:"spelling"`
	Items       []*Item   `json:"items"`
}

type Query struct {
	Request []*Request `json:"request"`
}

type Request struct {
	SearchTerms string `json:"searchTerms"`
	Count       int64  `json:"count"`
}

type Info struct {
	Time         float64 `json:"searchTime"`
	TotalResults string  `json:"totalResults"`
}

type Spelling struct {
	Corrected string `json:"correctedQuery"`
}

type Item struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	Snippet string `json:"snippet"`
}
