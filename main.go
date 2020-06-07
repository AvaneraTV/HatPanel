package main

import (
	"HatPanel/config"
	"fmt"
	"time"

	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/micmonay/keybd_event"
)

var (
	KB keybd_event.KeyBonding

	BS buttonState = buttonState{
		LastPressed:    -1,
		ButtonRegister: map[int]*widget.Button{},
	}
)

func main() {
	projectConfig := config.ReadConfig()

	a := app.New()
	a.Settings().SetTheme(&appTheme{a.Settings().Theme()})

	w := a.NewWindow("Hat Panel")
	layout.NewVBoxLayout()

	container := fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		// Row 1
		fyne.NewContainerWithLayout(layout.NewGridLayout(projectConfig.HotkeyPanel.NumCols),
			genButtons(projectConfig)...,
		),
	)

	w.SetContent(container)
	w.ShowAndRun()
}

func genButtons(c config.ProjectConfig) []fyne.CanvasObject {
	maxButtons := c.HotkeyPanel.NumCols * c.HotkeyPanel.NumRows
	buttons := []fyne.CanvasObject{}
	for i := 0; i < len(c.HotkeyPanel.Buttons) && i < maxButtons; i++ {
		var b fyne.CanvasObject
		b = newButton(c.HotkeyPanel.Buttons[i], i)
		buttons = append(buttons, b)
	}

	return buttons
}

func newButton(button config.HotkeyPanelButton, buttonConfigIndex int) *widget.Button {
	f := func() {
		KB.SetKeys(keybdEventFromKeyString(button.Key))
		KB.HasALT(button.HasAlt)
		KB.HasCTRL(button.HasCtrl)
		KB.HasSHIFT(button.HasShift)

		err := KB.Press()
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(10 * time.Millisecond)
		err = KB.Release()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(button.Key)
		}

		BS.LastPressed = buttonConfigIndex
		for k, v := range BS.ButtonRegister {
			if k == buttonConfigIndex {
				v.Style = widget.PrimaryButton
			} else {
				v.Style = widget.DefaultButton
			}
			v.Refresh()
		}
	}

	b := widget.NewButton(button.Label, f)
	BS.ButtonRegister[buttonConfigIndex] = b

	if button.IsDisabled {
		b.Hide()
	}

	return b
}

type buttonState struct {
	ButtonRegister map[int]*widget.Button
	LastPressed    int
}

type appTheme struct {
	defaultTheme fyne.Theme
}

func (t *appTheme) BackgroundColor() color.Color {
	return t.defaultTheme.BackgroundColor()
}
func (t *appTheme) ButtonColor() color.Color {
	return t.defaultTheme.ButtonColor()
}
func (t *appTheme) DisabledButtonColor() color.Color {
	return t.defaultTheme.DisabledButtonColor()
}
func (t *appTheme) HyperlinkColor() color.Color {
	return t.defaultTheme.HyperlinkColor()
}
func (t *appTheme) TextColor() color.Color {
	return t.defaultTheme.TextColor()
}
func (t *appTheme) DisabledTextColor() color.Color {
	return t.defaultTheme.DisabledTextColor()
}
func (t *appTheme) IconColor() color.Color {
	return t.defaultTheme.IconColor()
}
func (t *appTheme) DisabledIconColor() color.Color {
	return t.defaultTheme.DisabledIconColor()
}
func (t *appTheme) PlaceHolderColor() color.Color {
	return t.defaultTheme.PlaceHolderColor()
}
func (t *appTheme) PrimaryColor() color.Color {
	return t.defaultTheme.PrimaryColor()
}
func (t *appTheme) HoverColor() color.Color {
	return t.defaultTheme.HoverColor()
}
func (t *appTheme) FocusColor() color.Color {
	return t.defaultTheme.FocusColor()
}
func (t *appTheme) ScrollBarColor() color.Color {
	return t.defaultTheme.ScrollBarColor()
}
func (t *appTheme) ShadowColor() color.Color {
	return t.defaultTheme.ShadowColor()
}
func (t *appTheme) TextSize() int {
	return 32
	// t.defaultTheme.TextSize()
}
func (t *appTheme) TextFont() fyne.Resource {
	return t.defaultTheme.TextFont()
}
func (t *appTheme) TextBoldFont() fyne.Resource {
	return t.defaultTheme.TextBoldFont()
}
func (t *appTheme) TextItalicFont() fyne.Resource {
	return t.defaultTheme.TextItalicFont()
}
func (t *appTheme) TextBoldItalicFont() fyne.Resource {
	return t.defaultTheme.TextBoldItalicFont()
}
func (t *appTheme) TextMonospaceFont() fyne.Resource {
	return t.defaultTheme.TextMonospaceFont()
}
func (t *appTheme) Padding() int {
	return 32
	// return t.defaultTheme.Padding()
}
func (t *appTheme) IconInlineSize() int {
	return t.defaultTheme.IconInlineSize()
}
func (t *appTheme) ScrollBarSize() int {
	return t.defaultTheme.ScrollBarSize()
}
func (t *appTheme) ScrollBarSmallSize() int {
	return t.defaultTheme.ScrollBarSmallSize()
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
