package msconfig

import (
	"bufio"
	"io/ioutil"
	"os"
	"path"
	"runtime"

	"github.com/BurntSushi/toml"
)

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	} else if runtime.GOOS == "linux" {
		home := os.Getenv("XDG_CONFIG_HOME")
		if home != "" {
			return home
		}
	}
	return os.Getenv("HOME")
}

func configFilePath() string {
	return path.Join(userHomeDir(), ".manystagingsrc")
}

type ManyStagingsConfig struct {
	Token                string                 `toml:"token"`
	HostURL              string                 `toml:"host_url"`
	EnvironmentID        string                 `toml:"environment_id"`
	LogLevel             string                 `toml:"log_level"`
	OrchestratorProvider string                 `toml:"orchestrator_provider"`
	OrchestratorSettings map[string]interface{} `toml:"orchestrator_settings"`
}

func SaveConfig(config *ManyStagingsConfig) error {
	f, err := os.Create(configFilePath())
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	return toml.NewEncoder(w).Encode(&config)
}

func LoadConfig() (*ManyStagingsConfig, error) {
	var config ManyStagingsConfig
	configData, err := ioutil.ReadFile(configFilePath())
	if err != nil {
		return nil, err
	}
	if _, err := toml.Decode(string(configData), &config); err != nil {
		return nil, err
	}
	if config.LogLevel == "" {
		config.LogLevel = "info"
	}
	return &config, nil
}
