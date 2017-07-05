// +build solaris,!appengine

package logrus

import (
<<<<<<< HEAD
=======
	"io"
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	"os"

	"golang.org/x/sys/unix"
)

// IsTerminal returns true if the given file descriptor is a terminal.
<<<<<<< HEAD
func IsTerminal() bool {
	_, err := unix.IoctlGetTermios(int(os.Stdout.Fd()), unix.TCGETA)
	return err == nil
=======
func IsTerminal(f io.Writer) bool {
	switch v := f.(type) {
	case *os.File:
		_, err := unix.IoctlGetTermios(int(v.Fd()), unix.TCGETA)
		return err == nil
	default:
		return false
	}
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
}
