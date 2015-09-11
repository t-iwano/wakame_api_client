package apiconfig

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const ConfigFileName = "wakame.json"

var configDir = os.Getenv("WAKAME_CONFIG")

func homeDirGet() string {
	home := os.Getenv("HOME")
	return home
}

func init() {
	if configDir == "" {
		configDir = filepath.Join(homeDirGet(), ".wakame")
	}
}

func ConfigDir() string {
	return configDir
}

func SetConfigDir(dir string) {
	configDir = dir
}

type ConfigFile struct {
	wakame   *WakameConfig
	filename string
}

type WakameConfig struct {
	Scheme    string `schema,omitempty`
	Host      string `host,omitempty`
	Port      string `port,omitempty`
	Path      string `path,omitempty`
	AccountId string `accountid,omitempty`
}

func Load(configDir string) (*ConfigFile, error) {
	if configDir == "" {
		configDir = ConfigDir()
	}

	configFile := ConfigFile{
		wakame:   &WakameConfig{},
		filename: filepath.Join(configDir, ConfigFileName),
	}

	if _, err := os.Stat(configFile.filename); err == nil {
		file, err := os.Open(configFile.filename)
		if err != nil {
			return &configFile, err
		}
		defer file.Close()

		if err := json.NewDecoder(file).Decode(&configFile.wakame); err != nil {
			return &configFile, err
		}
		return &configFile, nil
	} else if !os.IsNotExist(err) {
		return &configFile, err
	}
	return &configFile, nil
}

func (configFile *ConfigFile) Wakame() *WakameConfig {
	return configFile.wakame
}

func (configFile *ConfigFile) Filename() string {
	return configFile.filename
}
