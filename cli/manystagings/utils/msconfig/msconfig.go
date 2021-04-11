package msconfig

import (
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

type ManyStagingsConfig struct {
	KubeconfigBase64 string `toml:"kubeconfig_base64"`
}

func LoadManyStagingsConfig() (*ManyStagingsConfig, error) {
	var config ManyStagingsConfig
	configData, err := ioutil.ReadFile(path.Join(userHomeDir(), ".manystagingsrc"))
	if err != nil {
		return nil, err
	}
	if _, err := toml.Decode(string(configData), &config); err != nil {
		return nil, err
	}
	return &config, nil
}
