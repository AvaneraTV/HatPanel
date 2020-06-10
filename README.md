# Installation

Executables are included for releases. A config file will be generated automatically when the program is run for the first time.

To build this project from source, the following requirements are needed:
- Go (tested using Go 1.14.3, may work with other versions)
- A gcc compiler. (TDM-GCC-64 has been tested)

The project can then be built using `go build` in the main directory.

# Configuration
App configuration is done in the application itself. Changes are real-time, but
are not saved to file until the user requests it. This allows the user to test
changes without accidently deleting their changes.

When the config file is saved by the program, it does not attempt to gracefully merge with the existing contents. Manual changes to the config file are discouraged unless care is taken.

# Disclaimer

This is a small hobby project, and hasn't been designed with performance, security, or backwards-compatibility (config will remain backwards compatible where possible) in mind. Please use at your own risk.