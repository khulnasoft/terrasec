//go:build windows
// +build windows

package cli

import (
	"os"

	"golang.org/x/sys/windows"
)

// In order for colored output to work on Windows, we need to explicitly
// enable the virtual terminal processing console mode.  This is needed for
// zap (log) output that goes to the console as well as program output.
func init() {
	var originalMode uint32

	stdout := windows.Handle(os.Stdout.Fd())

	windows.GetConsoleMode(stdout, &originalMode)
	windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}
