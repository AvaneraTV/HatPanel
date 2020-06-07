package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

/*
	This code exists to allow me to port config from the oldest version
	of the project. This kind of backwards compatibility isn't strictly
	necessary since the project is v0, however the author has distributed
	builds to a number of different non-technical individuals and would like
	to make the use of this tool as easy as possible.
*/

var (
	keys = [][]string{
		{"key1_label", "key1_keys"},
		{"key2_label", "key2_keys"},
		{"key3_label", "key3_keys"},
		{"key4_label", "key4_keys"},
		{"key5_label", "key5_keys"},
		{"key6_label", "key6_keys"},
		{"key7_label", "key7_keys"},
		{"key8_label", "key8_keys"},
		{"key9_label", "key9_keys"},
		{"key10_label", "key10_keys"},
		{"key11_label", "key1_1keys"}, // This line has a typo. Will fix after backwards-compatibility is added to config.
		{"key12_label", "key1_k2eys"}, // This line has a typo. Will fix after backwards-compatibility is added to config.
	}
)

const (
	legacyConfigFileName = "hatPanel_config.yml"
)

// readConfig_0_0_1 will read in the config that came with version 0.0.1 of the project.
// It is provided for backwards compatibility and no guarantees regarding maintenance are made.
// It should be safe to remove this once the early adopters have updated to the new config.
func readLegacy() (hadLegacy bool) {
	fmt.Println("Porting legacy config...")

	viper.SetConfigFile(legacyConfigFileName)
	err := viper.ReadInConfig()
	if err != nil {
		if strings.Contains(err.Error(), "The system cannot find the file specified") {
			return false
		}
	}

	viper.SetDefault("fragile_numrows", 3)
	viper.SetDefault("fragile_numcols", 4)

	newConfig := ProjectConfig{
		ConfigVersion: "0.1.0",
		HotkeyPanel: HotkeyPanelConfig{
			NumCols: viper.GetInt("fragile_numcols"),
			NumRows: viper.GetInt("fragile_numrows"),
			Buttons: []HotkeyPanelButton{},
		},
	}

	for i := range keys {
		b := HotkeyPanelButton{
			Label: viper.GetString(keys[i][0]),
		}
		workingKeyStr := viper.GetString(keys[i][1])
		b.HasCtrl = strings.Contains(workingKeyStr, "CTRL")
		workingKeyStr = strings.Replace(workingKeyStr, "CTRL+", "", -1)

		b.HasShift = strings.Contains(workingKeyStr, "SHIFT")
		workingKeyStr = strings.Replace(workingKeyStr, "SHIFT+", "", -1)

		b.HasAlt = strings.Contains(workingKeyStr, "ALT")
		workingKeyStr = strings.Replace(workingKeyStr, "ALT+", "", -1)

		b.Key = workingKeyStr
		b.IsDisabled = workingKeyStr == "DISABLED"

		newConfig.HotkeyPanel.Buttons = append(newConfig.HotkeyPanel.Buttons, b)
	}

	setAndWriteConfig(newConfig)
	ReadConfig()
	return true
}
