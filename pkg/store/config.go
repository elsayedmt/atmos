package store

import "github.com/mitchellh/mapstructure"

type StoreConfig struct {
	Type    string                 `yaml:"type"`
	Options map[string]interface{} `yaml:"options"`
}

type StoresConfig = map[string]StoreConfig

func parseOptions(options map[string]interface{}, target interface{}) error {
	return mapstructure.Decode(options, target)
}
