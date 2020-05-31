# Installation

To build this project from source, the following requirements are needed:
- Go (tested using Go 1.14.3, may work with other versions)
- A gcc compiler. (TDM-GCC-64 has been tested)

The project can be built using `go build` in the main directory. Or, by running the executable in the release.

# Configuration
Layout, hotkey management, etc. are all done manually via config file. This will be improved in future revisions. I will do everything I 
can to make config changes backwards-compatible and/or to automatically port config to avoid disruptions. In the far future, config management
will hopefully exist in the UI.

Config currently looks like this:
```yml
key1_keys: CTRL+U
key1_label: "1"
key2_keys: CTRL+I
key2_label: "2"
key3_keys: CTRL+O
key3_label: "3"
```

You can set hotkeys with any combination of CTRL/SHIFT/ALT modifiers, and a single other key. You don't need to use any modifiers at all if you don't want to. So, you could have "P", "CTRL+P", "CTRL+ALT+P", etc.

There is a known typo for the following 2 config labels which will be repaired in the future.
```yml
key1_1keys: CTRL+,
key1_k2eys: CTRL+.
```

# Disclaimer

This is a small hobby project, and hasn't been designed with performance, security, or backwards-compatibility in mind. Please use at your own risk.