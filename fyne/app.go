package fyne

import (
	"HatPanel/config"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
)

type App struct {
	fyneApp fyne.App

	hotkeyPanel *hotkeyPanel
}

// NewApp will set up and launch a new app. This app is the core of the HatPanel.
func NewApp() App {
	fyneApp := app.New()
	fyneApp.Settings().SetTheme(newAppTheme(fyneApp.Settings().Theme()))

	return App{
		fyneApp: fyneApp,
	}
}

// LaunchHotkeyPanel will launch the virtual keyboard panel with the most recently updated configuration.
func (a *App) LaunchHotkeyPanel(config config.HotkeyPanelConfig) {
	newWindow := a.fyneApp.NewWindow("HatPanel - Hotkeys")
	newWindow.SetOnClosed(a.OpenConfigPanel)
	hotkeyPanel := newHotkeyPanel(&newWindow, config)
	a.hotkeyPanel = &hotkeyPanel
}

// OpenConfigPanel will open the config panel. This is not finished, and this stub is a placeholder used for the hotkey panel refactor.
func (a *App) OpenConfigPanel() {
	a.fyneApp.Quit()
}
