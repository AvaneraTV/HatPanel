package main

import (
	"HatPanel/config"
	"HatPanel/keybd_event"

	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

var (
	// BS is used to avoid a deadlock wherin I need to allow myself to trigger actions on a button
	// when any button (including the same button) is pressed. Because it's self-referential, I need a secondary
	// reference to the button. It also allows me to reference the previous button pressed.
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
	keyboardEvent := keybd_event.GenerateKeypressFunction(button)
	f := func() {
		keyboardEvent()

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
