package utils

import (
	"os"

	"golang.org/x/sys/unix"
)

// IsTerminal checks if run in interactive shell
func IsTerminal() bool {
	_, err := unix.IoctlGetTermios(int(os.Stdout.Fd()), unix.TCGETS)

	return err == nil
}
