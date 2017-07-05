// +build appengine

package logrus

<<<<<<< HEAD
// IsTerminal returns true if stderr's file descriptor is a terminal.
func IsTerminal() bool {
=======
import "io"

// IsTerminal returns true if stderr's file descriptor is a terminal.
func IsTerminal(f io.Writer) bool {
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	return true
}
