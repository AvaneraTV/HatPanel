package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/micmonay/keybd_event"
	"github.com/spf13/viper"
)

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

	// Buttons are the parsed key/values from config.
	Buttons = []Button{}
)

func ReadConfig() {
	configFileName := "hatPanel_config.yml"

	viper.SetDefault(keys[0][0], "1")
	viper.SetDefault(keys[1][0], "2")
	viper.SetDefault(keys[2][0], "3")
	viper.SetDefault(keys[3][0], "4")
	viper.SetDefault(keys[4][0], "5")
	viper.SetDefault(keys[5][0], "6")
	viper.SetDefault(keys[6][0], "7")
	viper.SetDefault(keys[7][0], "8")
	viper.SetDefault(keys[8][0], "9")
	viper.SetDefault(keys[9][0], "10")
	viper.SetDefault(keys[10][0], "11")
	viper.SetDefault(keys[11][0], "12")

	viper.SetDefault(keys[0][1], "CTRL+U")
	viper.SetDefault(keys[1][1], "CTRL+I")
	viper.SetDefault(keys[2][1], "CTRL+O")
	viper.SetDefault(keys[3][1], "CTRL+P")
	viper.SetDefault(keys[4][1], "CTRL+J")
	viper.SetDefault(keys[5][1], "CTRL+K")
	viper.SetDefault(keys[6][1], "CTRL+L")
	viper.SetDefault(keys[7][1], "CTRL+;")
	viper.SetDefault(keys[8][1], "CTRL+N")
	viper.SetDefault(keys[9][1], "CTRL+M")
	viper.SetDefault(keys[10][1], "CTRL+,")
	viper.SetDefault(keys[11][1], "CTRL+.")

	viper.SetConfigFile(configFileName)
	err := viper.ReadInConfig()
	if err != nil {
		if strings.Contains(err.Error(), "The system cannot find the file specified") {
			viper.WriteConfigAs(configFileName)
			fmt.Println("Generated config file, please make any desired changes and re-launch. Program will exit in 10 seconds.")
			time.Sleep(10 * time.Second)
			os.Exit(0)
		} else {
			panic("Failed to read in config: " + err.Error())
		}
	}

	for i := range keys {
		b := Button{
			ID:        i,
			Label:     viper.GetString(keys[i][0]),
			keyString: viper.GetString(keys[i][1]),
		}
		b.parseKeyString()
	}
}

func Get(key string) string {
	return viper.GetString(key)
}

type Button struct {
	ID int

	keyString string

	Label string
	Key   int

	HasCtrl    bool
	HasShift   bool
	HasAlt     bool
	IsDisabled bool
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

	Buttons = append(Buttons, *b)
}

func (b *Button) String() string {
	return fmt.Sprintf("%s || %t %t %t", b.keyString, b.HasAlt, b.HasCtrl, b.HasShift)
}
