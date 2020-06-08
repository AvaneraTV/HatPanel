package fyne

import (
	"HatPanel/config"
	"HatPanel/keybd_event"
	"fmt"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type configPanel struct {
	config          *config.ProjectConfig
	window          *fyne.Window
	container       *fyne.Container
	openHotkeyPanel func()
	saveConfig      func()

	closeConfigOnHotkey bool
	IsOpen              bool
}

func newConfigPanel(window *fyne.Window, openHotkeyFunction func(), config *config.ProjectConfig) configPanel {

	panel := configPanel{
		config:          config,
		window:          window,
		openHotkeyPanel: openHotkeyFunction,
		saveConfig:      config.SaveToFile,
		IsOpen:          true,
	}
	panel.generateContainerContents()
	(*window).SetContent(panel.container)
	// panel.launch()
	return panel
}

func (cp *configPanel) launch() {
	(*cp.container).Refresh()
	(*cp.window).ShowAndRun()
}

func (cp *configPanel) show() {
	(*cp.container).Refresh()
	(*cp.window).Show()
}

func (cp *configPanel) refresh() {
	cp.generateContainerContents()
	(*cp.window).SetContent(cp.container)
}

func (cp *configPanel) generateContainerContents() {
	helpText := widget.NewLabel("Editing the config on this page will apply immediately, but isn't saved to file until you press that button.\n" +
		"So you can experiment without worrying about the effects. Closing the program will discard changes.")
	helpText.Alignment = fyne.TextAlignCenter
	cp.container = fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		cp.showTopButtons(),
		helpText,
		cp.showColRowEditor(),
		cp.generateButtonConfigRows(),
	)
}

func (cp *configPanel) showTopButtons() *fyne.Container {
	return fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		widget.NewButton("Open Hotkey Panel", cp.triggerHotkeyPress),
		widget.NewButton("Save config to file", cp.saveConfig),
	)
}

func (cp *configPanel) showColRowEditor() *fyne.Container {
	colEntry := widget.NewEntry()
	colEntry.SetText(strconv.Itoa(cp.config.HotkeyPanel.NumCols))
	colLabel := widget.NewLabel("Number of columns")
	colLabel.Alignment = fyne.TextAlignCenter
	colEntry.OnChanged = func(s string) {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Invalid entry for new number of columns: " + s)
		} else {
			cp.config.HotkeyPanel.SetNumCols(i)
		}
	}

	rowEntry := widget.NewEntry()
	rowEntry.SetText(strconv.Itoa(cp.config.HotkeyPanel.NumRows))
	rowLabel := widget.NewLabel("Number of rows")
	rowLabel.Alignment = fyne.TextAlignCenter
	rowEntry.OnChanged = func(s string) {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Invalid entry for new number of rows: " + s)
		} else {
			cp.config.HotkeyPanel.SetNumRows(i)
		}
	}

	return fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		colLabel,
		colEntry,
		rowLabel,
		rowEntry,
	)
}

func (cp *configPanel) generateButtonConfigRows() *fyne.Container {
	items := []fyne.CanvasObject{}

	for i := range cp.config.HotkeyPanel.Buttons {
		index := i

		movementButtons := fyne.NewContainerWithLayout(layout.NewGridLayout(2),
			widget.NewButton("^", func() {
				cp.config.HotkeyPanel.ShiftButtonUp(index)
				cp.refresh()
			}),
			widget.NewButton("v", func() {
				cp.config.HotkeyPanel.ShiftButtonDown(index)
				cp.refresh()
			}),
		)

		altCheck := widget.NewCheck("ALT", cp.config.HotkeyPanel.Buttons[index].SetAlt)
		altCheck.SetChecked(cp.config.HotkeyPanel.Buttons[index].HasAlt)
		ctrlCheck := widget.NewCheck("CTRL", cp.config.HotkeyPanel.Buttons[index].SetCtrl)
		ctrlCheck.SetChecked(cp.config.HotkeyPanel.Buttons[index].HasCtrl)
		shiftCheck := widget.NewCheck("SHIFT", cp.config.HotkeyPanel.Buttons[index].SetShift)
		shiftCheck.SetChecked(cp.config.HotkeyPanel.Buttons[index].HasShift)

		keySelect := widget.NewSelect(keybd_event.GetAllValidKeyStrings(), cp.config.HotkeyPanel.Buttons[index].SetKey)
		keySelect.SetSelected(cp.config.HotkeyPanel.Buttons[index].Key)

		label := widget.NewEntry()
		label.SetText(cp.config.HotkeyPanel.Buttons[index].Label)
		label.OnChanged = func(s string) {
			cp.config.HotkeyPanel.Buttons[index].SetLabel(s)
		}

		deleteBtn := widget.NewButton("Delete", func() {
			cp.config.HotkeyPanel.RemoveButton(index)
			cp.refresh()
		})

		items = append(items, movementButtons, altCheck, ctrlCheck, shiftCheck, keySelect, label, deleteBtn)
	}

	// Final button is to add a new row to the config.
	addButton := widget.NewButton("Add", func() {
		cp.config.HotkeyPanel.AddButton()
		cp.refresh()
	})
	items = append(items, addButton)

	c := fyne.NewContainerWithLayout(layout.NewGridLayout(7),
		items...,
	)

	return c
}

func (cp *configPanel) triggerHotkeyPress() {
	(*cp.window).Hide()
	cp.openHotkeyPanel()
}
