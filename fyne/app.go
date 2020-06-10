package fyne

import (
	"HatPanel/config"
	"HatPanel/util"
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
)

type App struct {
	fyneApp fyne.App

	config *config.ProjectConfig

	hotkeyPanel *hotkeyPanel
	configPanel *configPanel

	hotkeyTheme appTheme
	configTheme appTheme
}

// NewApp will set up and launch a new app. This app is the core of the HatPanel.
func NewApp(config *config.ProjectConfig) App {

	a := App{
		fyneApp: app.New(),
		config:  config,
	}
	a.configTheme = appTheme{
		defaultTheme: a.fyneApp.Settings().Theme(),
		textSize:     util.IntPtr(18),
		padding:      util.IntPtr(12),
	}
	a.hotkeyTheme = appTheme{
		defaultTheme: a.fyneApp.Settings().Theme(),
		textSize:     util.IntPtr(32),
		padding:      util.IntPtr(32),
	}

	a.initializeConfigPanel()

	return a
}

func (a *App) initializeConfigPanel() {
	a.fyneApp.Settings().SetTheme(&a.configTheme)
	showHotkeyPanel := func() {
		a.initializeHotkeyPanel()
		a.showHotkeyPanel()
	}

	configWindow := a.fyneApp.NewWindow("HatPanel - Config")
	configWindow.SetMaster()
	configPanel := newConfigPanel(&configWindow, showHotkeyPanel, a.config)
	a.configPanel = &configPanel

}

func (a *App) initializeHotkeyPanel() {
	a.fyneApp.Settings().SetTheme(&a.hotkeyTheme)

	hotkeyWindow := a.fyneApp.NewWindow("HatPanel - Hotkeys")
	hotkeyWindow.SetOnClosed(func() {
		a.initializeConfigPanel()
		a.configPanel.show()
	})
	hotkeyPanel := newHotkeyPanel(&hotkeyWindow, &a.config.HotkeyPanel)
	a.hotkeyPanel = &hotkeyPanel
}

func (a *App) Launch() {
	a.configPanel.launch()
}

// LaunchHotkeyPanel will launch the virtual keyboard panel with the most recently updated configuration.
func (a *App) showHotkeyPanel() {
	fmt.Println("Launching hotkey panel")
	a.hotkeyPanel.show()
}

// LaunchConfigPanel will open the config panel.
func (a *App) showConfigPanel() {
	fmt.Println("Launching config panel")
	a.configPanel.show()
}
