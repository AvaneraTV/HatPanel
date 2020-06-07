package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/micmonay/keybd_event"
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

var (
	// Buttons are the parsed key/values from config.
	Buttons = []Button{}
)

// ReadConfig will read in the configuration from file. It will save the config if needed.
func ReadConfig() ProjectConfig {

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
	return c
}

func setAndWriteConfig(c ProjectConfig) {
	v := viper.New()
	v.Set("config", c)
	v.SafeWriteConfigAs(configFileName)
}

type Button struct {
	ID int

	keyString string

	Label string
	Key   int

	HasCtrl      bool
	HasShift     bool
	HasAlt       bool
	IsDisabled   bool
	IsExceedsMax bool
}

func (b *Button) parseKeyString() {
	workingKeyStr := b.keyString
	b.HasCtrl = strings.Contains(workingKeyStr, "CTRL")
	workingKeyStr = strings.Replace(workingKeyStr, "CTRL+", "", -1)

	b.HasShift = strings.Contains(workingKeyStr, "SHIFT")
	workingKeyStr = strings.Replace(workingKeyStr, "SHIFT+", "", -1)

	b.HasAlt = strings.Contains(workingKeyStr, "ALT")
	workingKeyStr = strings.Replace(workingKeyStr, "ALT+", "", -1)

	switch workingKeyStr {
	case "~":
		b.Key = keybd_event.VK_SP1
	case "1":
		b.Key = keybd_event.VK_1
	case "2":
		b.Key = keybd_event.VK_2
	case "3":
		b.Key = keybd_event.VK_3
	case "4":
		b.Key = keybd_event.VK_4
	case "5":
		b.Key = keybd_event.VK_5
	case "6":
		b.Key = keybd_event.VK_6
	case "7":
		b.Key = keybd_event.VK_7
	case "8":
		b.Key = keybd_event.VK_8
	case "9":
		b.Key = keybd_event.VK_9
	case "0":
		b.Key = keybd_event.VK_0
	case "-":
		b.Key = keybd_event.VK_SP2
	case "=":
		b.Key = keybd_event.VK_SP3
	case "BACKSPACE":
		b.Key = keybd_event.VK_BACKSPACE
	case "TAB":
		b.Key = keybd_event.VK_TAB
	case "Q":
		b.Key = keybd_event.VK_Q
	case "W":
		b.Key = keybd_event.VK_W
	case "E":
		b.Key = keybd_event.VK_E
	case "R":
		b.Key = keybd_event.VK_R
	case "T":
		b.Key = keybd_event.VK_T
	case "Y":
		b.Key = keybd_event.VK_Y
	case "U":
		b.Key = keybd_event.VK_U
	case "I":
		b.Key = keybd_event.VK_I
	case "O":
		b.Key = keybd_event.VK_O
	case "P":
		b.Key = keybd_event.VK_P
	case "[":
		b.Key = keybd_event.VK_SP4
	case "]":
		b.Key = keybd_event.VK_SP5
	case "ENTER":
		b.Key = keybd_event.VK_ENTER
	case "CAPSLOCK":
		b.Key = keybd_event.VK_CAPSLOCK
	case "A":
		b.Key = keybd_event.VK_A
	case "S":
		b.Key = keybd_event.VK_S
	case "D":
		b.Key = keybd_event.VK_D
	case "F":
		b.Key = keybd_event.VK_F
	case "G":
		b.Key = keybd_event.VK_G
	case "H":
		b.Key = keybd_event.VK_H
	case "J":
		b.Key = keybd_event.VK_J
	case "K":
		b.Key = keybd_event.VK_K
	case "L":
		b.Key = keybd_event.VK_L
	case ";":
		b.Key = keybd_event.VK_SP6
	case "'":
		b.Key = keybd_event.VK_SP7
	case "\\":
		b.Key = keybd_event.VK_SP8
	case "Z":
		b.Key = keybd_event.VK_Z
	case "X":
		b.Key = keybd_event.VK_X
	case "C":
		b.Key = keybd_event.VK_C
	case "V":
		b.Key = keybd_event.VK_V
	case "B":
		b.Key = keybd_event.VK_B
	case "N":
		b.Key = keybd_event.VK_N
	case "M":
		b.Key = keybd_event.VK_M
	case ",":
		b.Key = keybd_event.VK_SP9
	case ".":
		b.Key = keybd_event.VK_SP10
	case "/":
		b.Key = keybd_event.VK_SP11
	case "SPACE":
		b.Key = keybd_event.VK_SPACE
	case "DISABLED":
		b.IsDisabled = true
	default:
		panic("don't know how to process " + b.keyString + ". Set the whole string to \"DISABLED\" if you'd like to ignore it.")
	}

	if !b.IsExceedsMax {
		Buttons = append(Buttons, *b)
	}
}

func (b *Button) String() string {
	return fmt.Sprintf("%s || %t %t %t", b.keyString, b.HasAlt, b.HasCtrl, b.HasShift)
}
