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
	Token                string `toml:"token"`
	EnvironmentID        string `toml:"environment_id"`
	OrchestratorProvider string `toml:"orchestrator_provider"`
	KubeconfigBase64     string `toml:"kubeconfig_base64"`
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
	return &config, nil
}
