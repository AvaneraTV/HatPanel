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
					Label: "1",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
				{
					Key:   "I",
					Label: "2",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
				{
					Key:   "O",
					Label: "3",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
				{
					Key:   "P",
					Label: "4",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
				{
					Key:   "J",
					Label: "5",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
				{
					Key:   "K",
					Label: "6",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
				{
					Key:   "L",
					Label: "7",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
				{
					Key:   ";",
					Label: "8",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
				{
					Key:   "N",
					Label: "9",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
				{
					Key:   "M",
					Label: "10",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
				{
					Key:   ",",
					Label: "11",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
				{
					Key:   ".",
					Label: "12",

					HasAlt:   false,
					HasCtrl:  true,
					HasShift: false,

					IsDisabled: false,
				},
			},
		},
	}
)
