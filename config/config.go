package config

import (
	"errors"
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
				defaultConfig.SaveToFile()
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

// SaveToFile will write the current config to the config file.
func (c *ProjectConfig) SaveToFile() {
	v := viper.New()
	v.Set("config", c)
	v.WriteConfigAs(configFileName)
	fmt.Println("Config saved to file")
}

// AddButton will add a new button to the config.
func (c *HotkeyPanelConfig) AddButton() {
	c.Buttons = append(c.Buttons, HotkeyPanelButton{
		Key:   "1",
		Label: "X",
	})
}

// RemoveButton will remove the specified button. Returns an error on out-of-bounds, or
func (c *HotkeyPanelConfig) RemoveButton(index int) error {
	if index < 0 || index >= len(c.Buttons) {
		return errors.New("invalid index for button removal")
	}

	c.Buttons = append(c.Buttons[:index], c.Buttons[index+1:]...)
	return nil
}

func (b *HotkeyPanelButton) String() string {
	s := ""
	if b.HasAlt {
		s += "ALT+"
	}
	if b.HasCtrl {
		s += "CTRL+"
	}
	if b.HasAlt {
		s += "ALT+"
	}
	s += b.Key

	return s
}

func (c *HotkeyPanelConfig) SetNumCols(new int) {
	c.NumCols = new
}

func (c *HotkeyPanelConfig) SetNumRows(new int) {
	c.NumRows = new
}

func (c *HotkeyPanelConfig) ShiftButtonUp(index int) {
	if index <= 0 {
		return
	}
	if index >= len(c.Buttons) {
		return
	}
	c.Buttons = append(append(c.Buttons[:index-1], c.Buttons[index], c.Buttons[index-1]), c.Buttons[index+1:]...)
}

func (c *HotkeyPanelConfig) ShiftButtonDown(index int) {
	if index >= len(c.Buttons)-1 {
		return
	}
	c.Buttons = append(append(c.Buttons[:index], c.Buttons[index+1], c.Buttons[index]), c.Buttons[index+2:]...)
}

func (c *HotkeyPanelButton) SetAlt(new bool) {
	c.HasAlt = new
}

func (c *HotkeyPanelButton) SetCtrl(new bool) {
	c.HasCtrl = new
}

func (c *HotkeyPanelButton) SetShift(new bool) {
	c.HasShift = new
}

func (c *HotkeyPanelButton) SetKey(new string) {
	c.Key = new
}

func (c *HotkeyPanelButton) SetLabel(new string) {
	c.Label = new
}
