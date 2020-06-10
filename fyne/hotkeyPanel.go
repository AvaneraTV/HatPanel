package fyne

import (
	"HatPanel/config"
	"HatPanel/keybd_event"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type hotkeyPanel struct {
	window      *fyne.Window
	container   *fyne.Container
	buttons     map[int]*widget.Button
	lastPressed int
}

type hotkeyPanelButtonState struct {
	ButtonRegister map[int]*widget.Button
	LastPressed    int
}

func newHotkeyPanel(window *fyne.Window, config *config.HotkeyPanelConfig) hotkeyPanel {
	panel := hotkeyPanel{
		window:  window,
		buttons: map[int]*widget.Button{},
	}

	container := fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		fyne.NewContainerWithLayout(layout.NewGridLayout(config.NumCols),
			panel.generateHotkeyButtons(config)...,
		),
	)
	panel.container = container
	(*window).SetContent(container)
	return panel
}

func (hp *hotkeyPanel) launch() {
	(*hp.container).Refresh()
	(*hp.window).ShowAndRun()
}

func (hp *hotkeyPanel) show() {
	(*hp.container).Refresh()
	(*hp.window).Show()
}

func (hp *hotkeyPanel) generateHotkeyButtons(c *config.HotkeyPanelConfig) []fyne.CanvasObject {
	maxButtons := c.NumCols * c.NumRows
	buttons := []fyne.CanvasObject{}
	for i := 0; i < len(c.Buttons) && i < maxButtons; i++ {
		var b fyne.CanvasObject
		b = hp.newHotkeyButton(&c.Buttons[i], i)
		buttons = append(buttons, b)
	}

	return buttons
}

func (hp *hotkeyPanel) newHotkeyButton(button *config.HotkeyPanelButton, buttonConfigIndex int) *widget.Button {
	keyboardEvent := keybd_event.GenerateKeypressFunction(*button)
	f := func() {
		keyboardEvent()

		hp.lastPressed = buttonConfigIndex
		for k, v := range hp.buttons {
			if k == buttonConfigIndex {
				v.Style = widget.PrimaryButton
			} else {
				v.Style = widget.DefaultButton
			}
			v.Refresh()
		}
	}

	b := widget.NewButton(button.Label, f)
	hp.buttons[buttonConfigIndex] = b

	if button.IsDisabled {
		b.Hide()
	}

	return b
}
