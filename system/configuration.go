package system

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

//Configs contains application configurations for all application modes
type Configs struct {
	Debug   Config
	Release Config
}

//Config contains application configuration for active application mode
type Config struct {
	Public   string `json:"public"`
	Domain   string `json:"domain"`
	Database DatabaseConfig
}

//DatabaseConfig contains database connection info
type DatabaseConfig struct {
	Host     string
	Name     string //database name
	User     string
	Password string
}

var (
	config *Config
)

//loadConfig unmarshals config for current application mode
func loadConfig() {
	data, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		panic(err)
	}
	configs := &Configs{}
	if err := json.Unmarshal(data, configs); err != nil {
		panic(err)
	}
	switch GetMode() {
	case ReleaseMode:
		config = &configs.Release
	default:
		config = &configs.Debug
	}
	if !path.IsAbs(config.Public) {
		workingDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		config.Public = path.Join(workingDir, config.Public)
	}
}

//GetConfig returns actual config
func GetConfig() *Config {
	return config
}
