package keybd_event

import (
	"HatPanel/config"
	"fmt"
	"sync"
	"time"

	"github.com/micmonay/keybd_event"
)

var (
	kb     keybd_event.KeyBonding
	kbLock sync.Mutex
)

// KeypressFunction is a type alias for a generic function, used to represent functions
// which result in a key-press with the virtual keyboard.
type KeypressFunction func()

// GenerateKeypressFunction will generate a KeypressFunction for the specfied button.
func GenerateKeypressFunction(button config.HotkeyPanelButton) KeypressFunction {
	return func() {
		// Mutex here is probably overkill, since this function should take only a few ms. But, someone somewhere
		// will do something weird if I'm not careful and I'll never figure out how to reproduce it.
		kbLock.Lock()
		defer kbLock.Unlock()

		kb.SetKeys(keybdEventFromKeyString(button.Key))
		kb.HasALT(button.HasAlt)
		kb.HasCTRL(button.HasCtrl)
		kb.HasSHIFT(button.HasShift)

		err := kb.Press()
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(10 * time.Millisecond)
		err = kb.Release()
		if err != nil {
			fmt.Println(err)
		} else {
			if button.HasAlt {
				fmt.Print("ALT+")
			}
			if button.HasCtrl {
				fmt.Print("CTRL+")
			}
			if button.HasAlt {
				fmt.Print("ALT+")
			}
			fmt.Println(button.Key)
		}
	}
}

func keybdEventFromKeyString(key string) int {
	switch key {
	case "~":
		return keybd_event.VK_SP1
	case "1":
		return keybd_event.VK_1
	case "2":
		return keybd_event.VK_2
	case "3":
		return keybd_event.VK_3
	case "4":
		return keybd_event.VK_4
	case "5":
		return keybd_event.VK_5
	case "6":
		return keybd_event.VK_6
	case "7":
		return keybd_event.VK_7
	case "8":
		return keybd_event.VK_8
	case "9":
		return keybd_event.VK_9
	case "0":
		return keybd_event.VK_0
	case "-":
		return keybd_event.VK_SP2
	case "=":
		return keybd_event.VK_SP3
	case "BACKSPACE":
		return keybd_event.VK_BACKSPACE
	case "TAB":
		return keybd_event.VK_TAB
	case "Q":
		return keybd_event.VK_Q
	case "W":
		return keybd_event.VK_W
	case "E":
		return keybd_event.VK_E
	case "R":
		return keybd_event.VK_R
	case "T":
		return keybd_event.VK_T
	case "Y":
		return keybd_event.VK_Y
	case "U":
		return keybd_event.VK_U
	case "I":
		return keybd_event.VK_I
	case "O":
		return keybd_event.VK_O
	case "P":
		return keybd_event.VK_P
	case "[":
		return keybd_event.VK_SP4
	case "]":
		return keybd_event.VK_SP5
	case "ENTER":
		return keybd_event.VK_ENTER
	case "CAPSLOCK":
		return keybd_event.VK_CAPSLOCK
	case "A":
		return keybd_event.VK_A
	case "S":
		return keybd_event.VK_S
	case "D":
		return keybd_event.VK_D
	case "F":
		return keybd_event.VK_F
	case "G":
		return keybd_event.VK_G
	case "H":
		return keybd_event.VK_H
	case "J":
		return keybd_event.VK_J
	case "K":
		return keybd_event.VK_K
	case "L":
		return keybd_event.VK_L
	case ";":
		return keybd_event.VK_SP6
	case "'":
		return keybd_event.VK_SP7
	case "\\":
		return keybd_event.VK_SP8
	case "Z":
		return keybd_event.VK_Z
	case "X":
		return keybd_event.VK_X
	case "C":
		return keybd_event.VK_C
	case "V":
		return keybd_event.VK_V
	case "B":
		return keybd_event.VK_B
	case "N":
		return keybd_event.VK_N
	case "M":
		return keybd_event.VK_M
	case ",":
		return keybd_event.VK_SP9
	case ".":
		return keybd_event.VK_SP10
	case "/":
		return keybd_event.VK_SP11
	case "SPACE":
		return keybd_event.VK_SPACE
	default:
		panic("Invalid key provided")
	}
}
