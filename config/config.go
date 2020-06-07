package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// ProjectConfig is a top-level wrapper for config
type ProjectConfig struct {
	ConfigVersion string
	HotkeyPanel   HotkeyPanelConfig
}

// HotkeyPanelConfig is the definition of our configuration options for the hotkey/button panel.
type HotkeyPanelConfig struct {
	NumCols int
	NumRows int
	Buttons []HotkeyPanelButton
}

// HotkeyPanelButton defines the options for a single button on the hotkey panel.
type HotkeyPanelButton struct {
	Key   string
	Label string

	// Key-combination modifiers
	HasAlt   bool
	HasCtrl  bool
	HasShift bool

	// Behavioral switches
	IsDisabled bool
}

const (
	configFileName       = "hatPanel_config.json"
	currentConfigVersion = "0.1.0"
)

// ReadConfig will read in the configuration from file. It will save the config if needed.
func ReadConfig() *ProjectConfig {

	v := viper.New()

	configFileName := configFileName
	v.SetConfigFile(configFileName)
	err := v.ReadInConfig()
	if err != nil {
		if strings.Contains(err.Error(), "The system cannot find the file specified") {
			hasLegacy := readLegacy()
			if !hasLegacy {
				// Write the default config
				setAndWriteConfig(defaultConfig)
				fmt.Println("Generated config file, please make any desired changes and re-launch. Program will exit in 10 seconds.")
				time.Sleep(10 * time.Second)
				os.Exit(0)
			}
		} else {
			panic("Failed to read in config: " + err.Error())
		}
	}

	c := ProjectConfig{}
	err = v.UnmarshalKey("config", &c)
	if err != nil {
		panic(err)
	}
	return &c
}

func setAndWriteConfig(c ProjectConfig) {
	v := viper.New()
	v.Set("config", c)
	v.SafeWriteConfigAs(configFileName)
}
