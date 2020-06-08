package config

var (
	defaultConfig = ProjectConfig{
		ConfigVersion: "0.1.0",
		HotkeyPanel: HotkeyPanelConfig{
			NumCols: 4,
			NumRows: 3,
			Buttons: []HotkeyPanelButton{
				{
					Key:   "U",
					Label: "Key1",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
				{
					Key:   "I",
					Label: "Key2",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
				{
					Key:   "O",
					Label: "Key3",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
				{
					Key:   "P",
					Label: "Key4",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
				{
					Key:   "J",
					Label: "Key5",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
				{
					Key:   "K",
					Label: "Key6",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
				{
					Key:   "L",
					Label: "Key7",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
				{
					Key:   ";",
					Label: "Key8",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
				{
					Key:   "N",
					Label: "Key9",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
				{
					Key:   "M",
					Label: "Key10",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
				{
					Key:   ",",
					Label: "Key11",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
				{
					Key:   ".",
					Label: "Key12",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
			},
		},
	}
)
