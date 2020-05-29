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
	config.ReadConfig()

	a := app.New()
	a.Settings().SetTheme(&appTheme{a.Settings().Theme()})

	w := a.NewWindow("Hat Panel")
	layout.NewVBoxLayout()

	container := fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		// Row 1
		fyne.NewContainerWithLayout(layout.NewGridLayout(2),
			newButton(config.Buttons[0]),
			newButton(config.Buttons[1]),
			newButton(config.Buttons[2]),
			newButton(config.Buttons[3]),
			newButton(config.Buttons[4]),
			newButton(config.Buttons[5]),
		),
	)

	w.SetContent(container)
	w.ShowAndRun()
}

func newButton(button config.Button) *widget.Button {
	f := func() {
		KB.SetKeys(button.Key)
		KB.HasALT(button.HasAlt)
		KB.HasCTRL(button.HasCtrl)
		KB.HasSHIFT(button.HasShift)
		// err := KB.Launching()
		// if err != nil {
		// 	fmt.Println(err)
		// } else {
		// 	fmt.Println(button.String())
		// }

		err := KB.Press()
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(10 * time.Millisecond)
		err = KB.Release()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(button.String())
		}

		BS.LastPressed = button.ID
		for k, v := range BS.ButtonRegister {
			if k == button.ID {
				v.Style = widget.PrimaryButton
			} else {
				v.Style = widget.DefaultButton
			}
			v.Refresh()
		}
	}

	b := widget.NewButton(button.Label, f)
	BS.ButtonRegister[button.ID] = b

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
